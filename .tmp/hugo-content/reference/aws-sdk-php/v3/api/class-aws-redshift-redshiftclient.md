Menu

- [Aws](namespace-aws.md)
- [Redshift](namespace-aws-redshift.md)

## RedshiftClient     extends [AwsClient](class-aws-awsclient.md)   in package    - [Aws](package-aws.md)

This client is used to interact with the **Amazon Redshift** service.

## Supported API Versions

This class uses a _service description model_ that is associated at
runtime based on the `version` option given when constructing the
client. The `version` option will determine which API operations,
waiters, and paginators are available for a client. Creating a command or a
specific API operation can be done using magic methods (e.g.,
`$client->commandName(/** parameters */)`, or using the
`$client->getCommand` method of the client.

- [**2012-12-01**](api-redshift-2012-12-01.md)

  - [AcceptReservedNodeExchange](api-redshift-2012-12-01-acceptreservednodeexchange.md)
  - [AddPartner](api-redshift-2012-12-01-addpartner.md)
  - [AssociateDataShareConsumer](api-redshift-2012-12-01-associatedatashareconsumer.md)
  - [AuthorizeClusterSecurityGroupIngress](api-redshift-2012-12-01-authorizeclustersecuritygroupingress.md)
  - [AuthorizeDataShare](api-redshift-2012-12-01-authorizedatashare.md)
  - [AuthorizeEndpointAccess](api-redshift-2012-12-01-authorizeendpointaccess.md)
  - [AuthorizeSnapshotAccess](api-redshift-2012-12-01-authorizesnapshotaccess.md)
  - [BatchDeleteClusterSnapshots](api-redshift-2012-12-01-batchdeleteclustersnapshots.md)
  - [BatchModifyClusterSnapshots](api-redshift-2012-12-01-batchmodifyclustersnapshots.md)
  - [CancelResize](api-redshift-2012-12-01-cancelresize.md)
  - [CopyClusterSnapshot](api-redshift-2012-12-01-copyclustersnapshot.md)
  - [CreateAuthenticationProfile](api-redshift-2012-12-01-createauthenticationprofile.md)
  - [CreateCluster](api-redshift-2012-12-01-createcluster.md)
  - [CreateClusterParameterGroup](api-redshift-2012-12-01-createclusterparametergroup.md)
  - [CreateClusterSecurityGroup](api-redshift-2012-12-01-createclustersecuritygroup.md)
  - [CreateClusterSnapshot](api-redshift-2012-12-01-createclustersnapshot.md)
  - [CreateClusterSubnetGroup](api-redshift-2012-12-01-createclustersubnetgroup.md)
  - [CreateCustomDomainAssociation](api-redshift-2012-12-01-createcustomdomainassociation.md)
  - [CreateEndpointAccess](api-redshift-2012-12-01-createendpointaccess.md)
  - [CreateEventSubscription](api-redshift-2012-12-01-createeventsubscription.md)
  - [CreateHsmClientCertificate](api-redshift-2012-12-01-createhsmclientcertificate.md)
  - [CreateHsmConfiguration](api-redshift-2012-12-01-createhsmconfiguration.md)
  - [CreateIntegration](api-redshift-2012-12-01-createintegration.md)
  - [CreateRedshiftIdcApplication](api-redshift-2012-12-01-createredshiftidcapplication.md)
  - [CreateScheduledAction](api-redshift-2012-12-01-createscheduledaction.md)
  - [CreateSnapshotCopyGrant](api-redshift-2012-12-01-createsnapshotcopygrant.md)
  - [CreateSnapshotSchedule](api-redshift-2012-12-01-createsnapshotschedule.md)
  - [CreateTags](api-redshift-2012-12-01-createtags.md)
  - [CreateUsageLimit](api-redshift-2012-12-01-createusagelimit.md)
  - [DeauthorizeDataShare](api-redshift-2012-12-01-deauthorizedatashare.md)
  - [DeleteAuthenticationProfile](api-redshift-2012-12-01-deleteauthenticationprofile.md)
  - [DeleteCluster](api-redshift-2012-12-01-deletecluster.md)
  - [DeleteClusterParameterGroup](api-redshift-2012-12-01-deleteclusterparametergroup.md)
  - [DeleteClusterSecurityGroup](api-redshift-2012-12-01-deleteclustersecuritygroup.md)
  - [DeleteClusterSnapshot](api-redshift-2012-12-01-deleteclustersnapshot.md)
  - [DeleteClusterSubnetGroup](api-redshift-2012-12-01-deleteclustersubnetgroup.md)
  - [DeleteCustomDomainAssociation](api-redshift-2012-12-01-deletecustomdomainassociation.md)
  - [DeleteEndpointAccess](api-redshift-2012-12-01-deleteendpointaccess.md)
  - [DeleteEventSubscription](api-redshift-2012-12-01-deleteeventsubscription.md)
  - [DeleteHsmClientCertificate](api-redshift-2012-12-01-deletehsmclientcertificate.md)
  - [DeleteHsmConfiguration](api-redshift-2012-12-01-deletehsmconfiguration.md)
  - [DeleteIntegration](api-redshift-2012-12-01-deleteintegration.md)
  - [DeletePartner](api-redshift-2012-12-01-deletepartner.md)
  - [DeleteRedshiftIdcApplication](api-redshift-2012-12-01-deleteredshiftidcapplication.md)
  - [DeleteResourcePolicy](api-redshift-2012-12-01-deleteresourcepolicy.md)
  - [DeleteScheduledAction](api-redshift-2012-12-01-deletescheduledaction.md)
  - [DeleteSnapshotCopyGrant](api-redshift-2012-12-01-deletesnapshotcopygrant.md)
  - [DeleteSnapshotSchedule](api-redshift-2012-12-01-deletesnapshotschedule.md)
  - [DeleteTags](api-redshift-2012-12-01-deletetags.md)
  - [DeleteUsageLimit](api-redshift-2012-12-01-deleteusagelimit.md)
  - [DeregisterNamespace](api-redshift-2012-12-01-deregisternamespace.md)
  - [DescribeAccountAttributes](api-redshift-2012-12-01-describeaccountattributes.md)
  - [DescribeAuthenticationProfiles](api-redshift-2012-12-01-describeauthenticationprofiles.md)
  - [DescribeClusterDbRevisions](api-redshift-2012-12-01-describeclusterdbrevisions.md)
  - [DescribeClusterParameterGroups](api-redshift-2012-12-01-describeclusterparametergroups.md)
  - [DescribeClusterParameters](api-redshift-2012-12-01-describeclusterparameters.md)
  - [DescribeClusterSecurityGroups](api-redshift-2012-12-01-describeclustersecuritygroups.md)
  - [DescribeClusterSnapshots](api-redshift-2012-12-01-describeclustersnapshots.md)
  - [DescribeClusterSubnetGroups](api-redshift-2012-12-01-describeclustersubnetgroups.md)
  - [DescribeClusterTracks](api-redshift-2012-12-01-describeclustertracks.md)
  - [DescribeClusterVersions](api-redshift-2012-12-01-describeclusterversions.md)
  - [DescribeClusters](api-redshift-2012-12-01-describeclusters.md)
  - [DescribeCustomDomainAssociations](api-redshift-2012-12-01-describecustomdomainassociations.md)
  - [DescribeDataShares](api-redshift-2012-12-01-describedatashares.md)
  - [DescribeDataSharesForConsumer](api-redshift-2012-12-01-describedatasharesforconsumer.md)
  - [DescribeDataSharesForProducer](api-redshift-2012-12-01-describedatasharesforproducer.md)
  - [DescribeDefaultClusterParameters](api-redshift-2012-12-01-describedefaultclusterparameters.md)
  - [DescribeEndpointAccess](api-redshift-2012-12-01-describeendpointaccess.md)
  - [DescribeEndpointAuthorization](api-redshift-2012-12-01-describeendpointauthorization.md)
  - [DescribeEventCategories](api-redshift-2012-12-01-describeeventcategories.md)
  - [DescribeEventSubscriptions](api-redshift-2012-12-01-describeeventsubscriptions.md)
  - [DescribeEvents](api-redshift-2012-12-01-describeevents.md)
  - [DescribeHsmClientCertificates](api-redshift-2012-12-01-describehsmclientcertificates.md)
  - [DescribeHsmConfigurations](api-redshift-2012-12-01-describehsmconfigurations.md)
  - [DescribeInboundIntegrations](api-redshift-2012-12-01-describeinboundintegrations.md)
  - [DescribeIntegrations](api-redshift-2012-12-01-describeintegrations.md)
  - [DescribeLoggingStatus](api-redshift-2012-12-01-describeloggingstatus.md)
  - [DescribeNodeConfigurationOptions](api-redshift-2012-12-01-describenodeconfigurationoptions.md)
  - [DescribeOrderableClusterOptions](api-redshift-2012-12-01-describeorderableclusteroptions.md)
  - [DescribePartners](api-redshift-2012-12-01-describepartners.md)
  - [DescribeRedshiftIdcApplications](api-redshift-2012-12-01-describeredshiftidcapplications.md)
  - [DescribeReservedNodeExchangeStatus](api-redshift-2012-12-01-describereservednodeexchangestatus.md)
  - [DescribeReservedNodeOfferings](api-redshift-2012-12-01-describereservednodeofferings.md)
  - [DescribeReservedNodes](api-redshift-2012-12-01-describereservednodes.md)
  - [DescribeResize](api-redshift-2012-12-01-describeresize.md)
  - [DescribeScheduledActions](api-redshift-2012-12-01-describescheduledactions.md)
  - [DescribeSnapshotCopyGrants](api-redshift-2012-12-01-describesnapshotcopygrants.md)
  - [DescribeSnapshotSchedules](api-redshift-2012-12-01-describesnapshotschedules.md)
  - [DescribeStorage](api-redshift-2012-12-01-describestorage.md)
  - [DescribeTableRestoreStatus](api-redshift-2012-12-01-describetablerestorestatus.md)
  - [DescribeTags](api-redshift-2012-12-01-describetags.md)
  - [DescribeUsageLimits](api-redshift-2012-12-01-describeusagelimits.md)
  - [DisableLogging](api-redshift-2012-12-01-disablelogging.md)
  - [DisableSnapshotCopy](api-redshift-2012-12-01-disablesnapshotcopy.md)
  - [DisassociateDataShareConsumer](api-redshift-2012-12-01-disassociatedatashareconsumer.md)
  - [EnableLogging](api-redshift-2012-12-01-enablelogging.md)
  - [EnableSnapshotCopy](api-redshift-2012-12-01-enablesnapshotcopy.md)
  - [FailoverPrimaryCompute](api-redshift-2012-12-01-failoverprimarycompute.md)
  - [GetClusterCredentials](api-redshift-2012-12-01-getclustercredentials.md)
  - [GetClusterCredentialsWithIAM](api-redshift-2012-12-01-getclustercredentialswithiam.md)
  - [GetIdentityCenterAuthToken](api-redshift-2012-12-01-getidentitycenterauthtoken.md)
  - [GetReservedNodeExchangeConfigurationOptions](api-redshift-2012-12-01-getreservednodeexchangeconfigurationoptions.md)
  - [GetReservedNodeExchangeOfferings](api-redshift-2012-12-01-getreservednodeexchangeofferings.md)
  - [GetResourcePolicy](api-redshift-2012-12-01-getresourcepolicy.md)
  - [ListRecommendations](api-redshift-2012-12-01-listrecommendations.md)
  - [ModifyAquaConfiguration](api-redshift-2012-12-01-modifyaquaconfiguration.md)
  - [ModifyAuthenticationProfile](api-redshift-2012-12-01-modifyauthenticationprofile.md)
  - [ModifyCluster](api-redshift-2012-12-01-modifycluster.md)
  - [ModifyClusterDbRevision](api-redshift-2012-12-01-modifyclusterdbrevision.md)
  - [ModifyClusterIamRoles](api-redshift-2012-12-01-modifyclusteriamroles.md)
  - [ModifyClusterMaintenance](api-redshift-2012-12-01-modifyclustermaintenance.md)
  - [ModifyClusterParameterGroup](api-redshift-2012-12-01-modifyclusterparametergroup.md)
  - [ModifyClusterSnapshot](api-redshift-2012-12-01-modifyclustersnapshot.md)
  - [ModifyClusterSnapshotSchedule](api-redshift-2012-12-01-modifyclustersnapshotschedule.md)
  - [ModifyClusterSubnetGroup](api-redshift-2012-12-01-modifyclustersubnetgroup.md)
  - [ModifyCustomDomainAssociation](api-redshift-2012-12-01-modifycustomdomainassociation.md)
  - [ModifyEndpointAccess](api-redshift-2012-12-01-modifyendpointaccess.md)
  - [ModifyEventSubscription](api-redshift-2012-12-01-modifyeventsubscription.md)
  - [ModifyIntegration](api-redshift-2012-12-01-modifyintegration.md)
  - [ModifyLakehouseConfiguration](api-redshift-2012-12-01-modifylakehouseconfiguration.md)
  - [ModifyRedshiftIdcApplication](api-redshift-2012-12-01-modifyredshiftidcapplication.md)
  - [ModifyScheduledAction](api-redshift-2012-12-01-modifyscheduledaction.md)
  - [ModifySnapshotCopyRetentionPeriod](api-redshift-2012-12-01-modifysnapshotcopyretentionperiod.md)
  - [ModifySnapshotSchedule](api-redshift-2012-12-01-modifysnapshotschedule.md)
  - [ModifyUsageLimit](api-redshift-2012-12-01-modifyusagelimit.md)
  - [PauseCluster](api-redshift-2012-12-01-pausecluster.md)
  - [PurchaseReservedNodeOffering](api-redshift-2012-12-01-purchasereservednodeoffering.md)
  - [PutResourcePolicy](api-redshift-2012-12-01-putresourcepolicy.md)
  - [RebootCluster](api-redshift-2012-12-01-rebootcluster.md)
  - [RegisterNamespace](api-redshift-2012-12-01-registernamespace.md)
  - [RejectDataShare](api-redshift-2012-12-01-rejectdatashare.md)
  - [ResetClusterParameterGroup](api-redshift-2012-12-01-resetclusterparametergroup.md)
  - [ResizeCluster](api-redshift-2012-12-01-resizecluster.md)
  - [RestoreFromClusterSnapshot](api-redshift-2012-12-01-restorefromclustersnapshot.md)
  - [RestoreTableFromClusterSnapshot](api-redshift-2012-12-01-restoretablefromclustersnapshot.md)
  - [ResumeCluster](api-redshift-2012-12-01-resumecluster.md)
  - [RevokeClusterSecurityGroupIngress](api-redshift-2012-12-01-revokeclustersecuritygroupingress.md)
  - [RevokeEndpointAccess](api-redshift-2012-12-01-revokeendpointaccess.md)
  - [RevokeSnapshotAccess](api-redshift-2012-12-01-revokesnapshotaccess.md)
  - [RotateEncryptionKey](api-redshift-2012-12-01-rotateencryptionkey.md)
  - [UpdatePartnerStatus](api-redshift-2012-12-01-updatepartnerstatus.md)

### Table of Contents  [header link](class-aws-redshift-redshiftclient-toc.md)

#### Methods  [header link](class-aws-redshift-redshiftclient-toc-methods.md)

[\_\_call()](class-aws-awsclienttrait.md#method___call)
: mixed [\_\_construct()](class-aws-awsclient.md#method___construct)
: mixed The client constructor accepts the following options:[\_\_sleep()](class-aws-awsclient.md#method___sleep)
: mixed [execute()](class-aws-awsclienttrait.md#method_execute)
: mixed [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
: mixed [factory()](class-aws-awsclient.md#method_factory)
: static [getApi()](class-aws-awsclienttrait.md#method_getApi)
: [Service](class-aws-api-service.md)[getArguments()](class-aws-awsclient.md#method_getArguments)
: array<string\|int, mixed> Get an array of client constructor arguments used by the client.[getClientBuiltIns()](class-aws-awsclient.md#method_getClientBuiltIns)
: array<string\|int, mixed> Provides the set of built-in keys and values
used for endpoint resolution[getClientContextParams()](class-aws-awsclient.md#method_getClientContextParams)
: array<string\|int, mixed> Provides the set of service context parameter
key-value pairs used for endpoint resolution.[getCommand()](class-aws-awsclienttrait.md#method_getCommand)
: [CommandInterface](class-aws-commandinterface.md)[getConfig()](class-aws-awsclient.md#method_getConfig)
: mixed\|null Get a client configuration value.[getCredentials()](class-aws-awsclient.md#method_getCredentials)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.[getEndpoint()](class-aws-awsclient.md#method_getEndpoint)
: [UriInterface](class-psr-http-message-uriinterface.md)Gets the default endpoint, or base URL, used by the client.[getEndpointProvider()](class-aws-awsclient.md#method_getEndpointProvider)
: mixed [getEndpointProviderArgs()](class-aws-awsclient.md#method_getEndpointProviderArgs)
: array<string\|int, mixed> Retrieves arguments to be used in endpoint resolution.[getHandlerList()](class-aws-awsclient.md#method_getHandlerList)
: [HandlerList](class-aws-handlerlist.md)Get the handler list used to transfer commands.[getIterator()](class-aws-awsclienttrait.md#method_getIterator)
: mixed [getPaginator()](class-aws-awsclienttrait.md#method_getPaginator)
: mixed [getRegion()](class-aws-awsclient.md#method_getRegion)
: string Get the region to which the client is configured to send requests.[getSignatureProvider()](class-aws-awsclient.md#method_getSignatureProvider)
: callable Get the signature\_provider function of the client.[getToken()](class-aws-awsclient.md#method_getToken)
: mixed [getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
: mixed [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)
: mixed

### Methods  [header link](class-aws-redshift-redshiftclient-methods.md)

#### \_\_call()  [header link](class-aws-awsclienttrait.md\#method___call)

`
    public
                    __call(mixed $name, array<string|int, mixed> $args) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>

#### \_\_construct()  [header link](class-aws-awsclient.md\#method___construct)

The client constructor accepts the following options:

`
    public
                    __construct(array<string|int, mixed> $args) : mixed`

- api\_provider: (callable) An optional PHP callable that accepts a
type, service, and version argument, and returns an array of
corresponding configuration data. The type value can be one of api,
waiter, or paginator.
- credentials:
(Aws\\Credentials\\CredentialsInterface\|array\|bool\|callable) Specifies
the credentials used to sign requests. Provide an
Aws\\Credentials\\CredentialsInterface object, an associative array of
"key", "secret", and an optional "token" key, `false` to use null
credentials, or a callable credentials provider used to create
credentials or return null. See Aws\\Credentials\\CredentialProvider for
a list of built-in credentials providers. If no credentials are
provided, the SDK will attempt to load them from the environment.
- token:
(Aws\\Token\\TokenInterface\|array\|bool\|callable) Specifies
the token used to authorize requests. Provide an
Aws\\Token\\TokenInterface object, an associative array of
"token" and an optional "expires" key, `false` to use no
token, or a callable token provider used to create a
token or return null. See Aws\\Token\\TokenProvider for
a list of built-in token providers. If no token is
provided, the SDK will attempt to load one from the environment.
- csm:
(Aws\\ClientSideMonitoring\\ConfigurationInterface\|array\|callable) Specifies
the credentials used to sign requests. Provide an
Aws\\ClientSideMonitoring\\ConfigurationInterface object, a callable
configuration provider used to create client-side monitoring configuration,
`false` to disable csm, or an associative array with the following keys:
enabled: (bool) Set to true to enable client-side monitoring, defaults
to false; host: (string) the host location to send monitoring events to,
defaults to 127.0.0.1; port: (int) The port used for the host connection,
defaults to 31000; client\_id: (string) An identifier for this project
- debug: (bool\|array) Set to true to display debug information when
sending requests. Alternatively, you can provide an associative array
with the following keys: logfn: (callable) Function that is invoked
with log messages; stream\_size: (int) When the size of a stream is
greater than this number, the stream data will not be logged (set to
"0" to not log any stream data); scrub\_auth: (bool) Set to false to
disable the scrubbing of auth data from the logged messages; http:
(bool) Set to false to disable the "debug" feature of lower level HTTP
adapters (e.g., verbose curl output).
- stats: (bool\|array) Set to true to gather transfer statistics on
requests sent. Alternatively, you can provide an associative array with
the following keys: retries: (bool) Set to false to disable reporting
on retries attempted; http: (bool) Set to true to enable collecting
statistics from lower level HTTP adapters (e.g., values returned in
GuzzleHttp\\TransferStats). HTTP handlers must support an
`http_stats_receiver` option for this to have an effect; timer: (bool)
Set to true to enable a command timer that reports the total wall clock
time spent on an operation in seconds.
- disable\_host\_prefix\_injection: (bool) Set to true to disable host prefix
injection logic for services that use it. This disables the entire
prefix injection, including the portions supplied by user-defined
parameters. Setting this flag will have no effect on services that do
not use host prefix injection.
- endpoint: (string) The full URI of the webservice. This is only
required when connecting to a custom endpoint (e.g., a local version
of S3).
- endpoint\_discovery: (Aws\\EndpointDiscovery\\ConfigurationInterface,
Aws\\CacheInterface, array, callable) Settings for endpoint discovery.
Provide an instance of Aws\\EndpointDiscovery\\ConfigurationInterface,
an instance Aws\\CacheInterface, a callable that provides a promise for
a Configuration object, or an associative array with the following
keys: enabled: (bool) Set to true to enable endpoint discovery, false
to explicitly disable it, defaults to false; cache\_limit: (int) The
maximum number of keys in the endpoints cache, defaults to 1000.
- endpoint\_provider: (callable) An optional PHP callable that
accepts a hash of options including a "service" and "region" key and
returns NULL or a hash of endpoint data, of which the "endpoint" key
is required. See Aws\\Endpoint\\EndpointProvider for a list of built-in
providers.
- handler: (callable) A handler that accepts a command object,
request object and returns a promise that is fulfilled with an
Aws\\ResultInterface object or rejected with an
Aws\\Exception\\AwsException. A handler does not accept a next handler
as it is terminal and expected to fulfill a command. If no handler is
provided, a default Guzzle handler will be utilized.
- http: (array, default=array(0)) Set to an array of SDK request
options to apply to each request (e.g., proxy, verify, etc.).
- http\_handler: (callable) An HTTP handler is a function that
accepts a PSR-7 request object and returns a promise that is fulfilled
with a PSR-7 response object or rejected with an array of exception
data. NOTE: This option supersedes any provided "handler" option.
- idempotency\_auto\_fill: (bool\|callable) Set to false to disable SDK to
populate parameters that enabled 'idempotencyToken' trait with a random
UUID v4 value on your behalf. Using default value 'true' still allows
parameter value to be overwritten when provided. Note: auto-fill only
works when cryptographically secure random bytes generator functions
(random\_bytes, openssl\_random\_pseudo\_bytes or mcrypt\_create\_iv) can be
found. You may also provide a callable source of random bytes.
- profile: (string) Allows you to specify which profile to use when
credentials are created from the AWS credentials file in your HOME
directory. This setting overrides the AWS\_PROFILE environment
variable. Note: Specifying "profile" will cause the "credentials" key
to be ignored.
- region: (string, required) Region to connect to. See
http://docs.aws.amazon.com/general/latest/gr/rande.html for a list of
available regions.
- retries: (int, Aws\\Retry\\ConfigurationInterface, Aws\\CacheInterface,
array, callable) Configures the retry mode and maximum number of
allowed retries for a client (pass 0 to disable retries). Provide an
integer for 'legacy' mode with the specified number of retries.
Otherwise provide an instance of Aws\\Retry\\ConfigurationInterface, an
instance of Aws\\CacheInterface, a callable function, or an array with
the following keys: mode: (string) Set to 'legacy', 'standard' (uses
retry quota management), or 'adapative' (an experimental mode that adds
client-side rate limiting to standard mode); max\_attempts (int) The
maximum number of attempts for a given request.
- scheme: (string, default=string(5) "https") URI scheme to use when
connecting connect. The SDK will utilize "https" endpoints (i.e.,
utilize SSL/TLS connections) by default. You can attempt to connect to
a service over an unencrypted "http" endpoint by setting `scheme` to
"http".
- signature\_provider: (callable) A callable that accepts a signature
version name (e.g., "v4"), a service name, and region, and
returns a SignatureInterface object or null. This provider is used to
create signers utilized by the client. See
Aws\\Signature\\SignatureProvider for a list of built-in providers
- signature\_version: (string) A string representing a custom
signature version to use with a service (e.g., v4). Note that
per/operation signature version MAY override this requested signature
version.
- use\_aws\_shared\_config\_files: (bool, default=bool(true)) Set to false to
disable checking for shared config file in '~/.aws/config' and
'~/.aws/credentials'. This will override the AWS\_CONFIG\_FILE
environment variable.
- validate: (bool, default=bool(true)) Set to false to disable
client-side parameter validation.
- version: (string, required) The version of the webservice to
utilize (e.g., 2006-03-01).
- account\_id\_endpoint\_mode: (string, default(preferred)) this option
decides whether credentials should resolve an accountId value,
which is going to be used as part of the endpoint resolution.
The valid values for this option are:
  - preferred: when this value is set then, a warning is logged when
    accountId is empty in the resolved identity.
  - required: when this value is set then, an exception is thrown when
    accountId is empty in the resolved identity.
  - disabled: when this value is set then, the validation for if accountId
    was resolved or not, is ignored.
- ua\_append: (string, array) To pass custom user agent parameters.
- app\_id: (string) an optional application specific identifier that can be set.
When set it will be appended to the User-Agent header of every request
in the form of App/{AppId}. This variable is sourced from environment
variable AWS\_SDK\_UA\_APP\_ID or the shared config profile attribute sdk\_ua\_app\_id.
See https://docs.aws.amazon.com/sdkref/latest/guide/settings-reference.html for
more information on environment variables and shared config settings.

##### Parameters

$args
: array<string\|int, mixed>

Client configuration arguments.

##### Tags  [header link](class-aws-awsclient.md\#method___construct\#tags)

throwsInvalidArgumentException

if any required options are missing or
the service is not supported.

#### \_\_sleep()  [header link](class-aws-awsclient.md\#method___sleep)

`
    public
                    __sleep() : mixed`

#### execute()  [header link](class-aws-awsclienttrait.md\#method_execute)

`
    public
                    execute(CommandInterface $command) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

#### executeAsync()  [header link](class-aws-awsclienttrait.md\#method_executeAsync)

`
    public
                    executeAsync(CommandInterface $command) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

#### factory()  [header link](class-aws-awsclient.md\#method_factory)

`
    public
            static        factory([array<string|int, mixed> $config = [] ]) : static`

##### Parameters

$config
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-awsclient.md\#method_factory\#tags)

deprecated

##### Return values

static

#### getApi()  [header link](class-aws-awsclienttrait.md\#method_getApi)

`
    public
    abstract                getApi() : Service`

##### Return values

[Service](class-aws-api-service.md)

#### getArguments()  [header link](class-aws-awsclient.md\#method_getArguments)

Get an array of client constructor arguments used by the client.

`
    public
            static        getArguments() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getClientBuiltIns()  [header link](class-aws-awsclient.md\#method_getClientBuiltIns)

Provides the set of built-in keys and values
used for endpoint resolution

`
    public
                    getClientBuiltIns() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getClientContextParams()  [header link](class-aws-awsclient.md\#method_getClientContextParams)

Provides the set of service context parameter
key-value pairs used for endpoint resolution.

`
    public
                    getClientContextParams() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getCommand()  [header link](class-aws-awsclienttrait.md\#method_getCommand)

`
    public
    abstract                getCommand(string $name[, array<string|int, mixed> $args = [] ]) : CommandInterface`

##### Parameters

$name
: string$args
: array<string\|int, mixed>
= \[\]

##### Return values

[CommandInterface](class-aws-commandinterface.md)

#### getConfig()  [header link](class-aws-awsclient.md\#method_getConfig)

Get a client configuration value.

`
    public
                    getConfig([mixed $option = null ]) : mixed|null`

##### Parameters

$option
: mixed
= null

The option to retrieve. Pass null to retrieve
all options.

##### Return values

mixed\|null

#### getCredentials()  [header link](class-aws-awsclient.md\#method_getCredentials)

Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.

`
    public
                    getCredentials() : PromiseInterface`

If you need the credentials synchronously, then call the wait() method
on the returned promise.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### getEndpoint()  [header link](class-aws-awsclient.md\#method_getEndpoint)

Gets the default endpoint, or base URL, used by the client.

`
    public
                    getEndpoint() : UriInterface`

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)

#### getEndpointProvider()  [header link](class-aws-awsclient.md\#method_getEndpointProvider)

`
    public
                    getEndpointProvider() : mixed`

#### getEndpointProviderArgs()  [header link](class-aws-awsclient.md\#method_getEndpointProviderArgs)

Retrieves arguments to be used in endpoint resolution.

`
    public
                    getEndpointProviderArgs() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getHandlerList()  [header link](class-aws-awsclient.md\#method_getHandlerList)

Get the handler list used to transfer commands.

`
    public
                    getHandlerList() : HandlerList`

This list can be modified to add middleware or to change the underlying
handler used to send HTTP requests.

##### Return values

[HandlerList](class-aws-handlerlist.md)

#### getIterator()  [header link](class-aws-awsclienttrait.md\#method_getIterator)

`
    public
                    getIterator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getPaginator()  [header link](class-aws-awsclienttrait.md\#method_getPaginator)

`
    public
                    getPaginator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getRegion()  [header link](class-aws-awsclient.md\#method_getRegion)

Get the region to which the client is configured to send requests.

`
    public
                    getRegion() : string`

##### Return values

string

#### getSignatureProvider()  [header link](class-aws-awsclient.md\#method_getSignatureProvider)

Get the signature\_provider function of the client.

`
    public
        final            getSignatureProvider() : callable`

##### Return values

callable

#### getToken()  [header link](class-aws-awsclient.md\#method_getToken)

`
    public
                    getToken() : mixed`

#### getWaiter()  [header link](class-aws-awsclienttrait.md\#method_getWaiter)

`
    public
                    getWaiter(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### waitUntil()  [header link](class-aws-awsclienttrait.md\#method_waitUntil)

`
    public
                    waitUntil(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-redshift-redshiftclient-toc-methods.md)
- Methods
  - [\_\_call()](class-aws-awsclienttrait.md#method___call)
  - [\_\_construct()](class-aws-awsclient.md#method___construct)
  - [\_\_sleep()](class-aws-awsclient.md#method___sleep)
  - [execute()](class-aws-awsclienttrait.md#method_execute)
  - [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
  - [factory()](class-aws-awsclient.md#method_factory)
  - [getApi()](class-aws-awsclienttrait.md#method_getApi)
  - [getArguments()](class-aws-awsclient.md#method_getArguments)
  - [getClientBuiltIns()](class-aws-awsclient.md#method_getClientBuiltIns)
  - [getClientContextParams()](class-aws-awsclient.md#method_getClientContextParams)
  - [getCommand()](class-aws-awsclienttrait.md#method_getCommand)
  - [getConfig()](class-aws-awsclient.md#method_getConfig)
  - [getCredentials()](class-aws-awsclient.md#method_getCredentials)
  - [getEndpoint()](class-aws-awsclient.md#method_getEndpoint)
  - [getEndpointProvider()](class-aws-awsclient.md#method_getEndpointProvider)
  - [getEndpointProviderArgs()](class-aws-awsclient.md#method_getEndpointProviderArgs)
  - [getHandlerList()](class-aws-awsclient.md#method_getHandlerList)
  - [getIterator()](class-aws-awsclienttrait.md#method_getIterator)
  - [getPaginator()](class-aws-awsclienttrait.md#method_getPaginator)
  - [getRegion()](class-aws-awsclient.md#method_getRegion)
  - [getSignatureProvider()](class-aws-awsclient.md#method_getSignatureProvider)
  - [getToken()](class-aws-awsclient.md#method_getToken)
  - [getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
  - [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)

[Back To Top](class-aws-redshift-redshiftclient-top.md)
