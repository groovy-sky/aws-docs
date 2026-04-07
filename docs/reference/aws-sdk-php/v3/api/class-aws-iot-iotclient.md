Menu

- [Aws](namespace-aws.md)
- [Iot](namespace-aws-iot.md)

## IotClient     extends [AwsClient](class-aws-awsclient.md)   in package    - [Aws](package-aws.md)

This client is used to interact with the **AWS IoT** service.

## Supported API Versions

This class uses a _service description model_ that is associated at
runtime based on the `version` option given when constructing the
client. The `version` option will determine which API operations,
waiters, and paginators are available for a client. Creating a command or a
specific API operation can be done using magic methods (e.g.,
`$client->commandName(/** parameters */)`, or using the
`$client->getCommand` method of the client.

- [**2015-05-28**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html)

  - [AcceptCertificateTransfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#acceptcertificatetransfer)
  - [AddThingToBillingGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#addthingtobillinggroup)
  - [AddThingToThingGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#addthingtothinggroup)
  - [AssociateSbomWithPackageVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#associatesbomwithpackageversion)
  - [AssociateTargetsWithJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#associatetargetswithjob)
  - [AttachPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#attachpolicy)
  - [AttachPrincipalPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#attachprincipalpolicy)
  - [AttachSecurityProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#attachsecurityprofile)
  - [AttachThingPrincipal](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#attachthingprincipal)
  - [CancelAuditMitigationActionsTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#cancelauditmitigationactionstask)
  - [CancelAuditTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#cancelaudittask)
  - [CancelCertificateTransfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#cancelcertificatetransfer)
  - [CancelDetectMitigationActionsTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#canceldetectmitigationactionstask)
  - [CancelJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#canceljob)
  - [CancelJobExecution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#canceljobexecution)
  - [ClearDefaultAuthorizer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#cleardefaultauthorizer)
  - [ConfirmTopicRuleDestination](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#confirmtopicruledestination)
  - [CreateAuditSuppression](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createauditsuppression)
  - [CreateAuthorizer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createauthorizer)
  - [CreateBillingGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createbillinggroup)
  - [CreateCertificateFromCsr](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createcertificatefromcsr)
  - [CreateCertificateProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createcertificateprovider)
  - [CreateCommand](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createcommand)
  - [CreateCustomMetric](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createcustommetric)
  - [CreateDimension](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createdimension)
  - [CreateDomainConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createdomainconfiguration)
  - [CreateDynamicThingGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createdynamicthinggroup)
  - [CreateFleetMetric](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createfleetmetric)
  - [CreateJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createjob)
  - [CreateJobTemplate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createjobtemplate)
  - [CreateKeysAndCertificate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createkeysandcertificate)
  - [CreateMitigationAction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createmitigationaction)
  - [CreateOTAUpdate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createotaupdate)
  - [CreatePackage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createpackage)
  - [CreatePackageVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createpackageversion)
  - [CreatePolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createpolicy)
  - [CreatePolicyVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createpolicyversion)
  - [CreateProvisioningClaim](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createprovisioningclaim)
  - [CreateProvisioningTemplate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createprovisioningtemplate)
  - [CreateProvisioningTemplateVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createprovisioningtemplateversion)
  - [CreateRoleAlias](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createrolealias)
  - [CreateScheduledAudit](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createscheduledaudit)
  - [CreateSecurityProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createsecurityprofile)
  - [CreateStream](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createstream)
  - [CreateThing](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#creatething)
  - [CreateThingGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createthinggroup)
  - [CreateThingType](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createthingtype)
  - [CreateTopicRule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createtopicrule)
  - [CreateTopicRuleDestination](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#createtopicruledestination)
  - [DeleteAccountAuditConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deleteaccountauditconfiguration)
  - [DeleteAuditSuppression](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deleteauditsuppression)
  - [DeleteAuthorizer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deleteauthorizer)
  - [DeleteBillingGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletebillinggroup)
  - [DeleteCACertificate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletecacertificate)
  - [DeleteCertificate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletecertificate)
  - [DeleteCertificateProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletecertificateprovider)
  - [DeleteCommand](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletecommand)
  - [DeleteCommandExecution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletecommandexecution)
  - [DeleteCustomMetric](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletecustommetric)
  - [DeleteDimension](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletedimension)
  - [DeleteDomainConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletedomainconfiguration)
  - [DeleteDynamicThingGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletedynamicthinggroup)
  - [DeleteFleetMetric](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletefleetmetric)
  - [DeleteJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletejob)
  - [DeleteJobExecution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletejobexecution)
  - [DeleteJobTemplate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletejobtemplate)
  - [DeleteMitigationAction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletemitigationaction)
  - [DeleteOTAUpdate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deleteotaupdate)
  - [DeletePackage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletepackage)
  - [DeletePackageVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletepackageversion)
  - [DeletePolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletepolicy)
  - [DeletePolicyVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletepolicyversion)
  - [DeleteProvisioningTemplate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deleteprovisioningtemplate)
  - [DeleteProvisioningTemplateVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deleteprovisioningtemplateversion)
  - [DeleteRegistrationCode](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deleteregistrationcode)
  - [DeleteRoleAlias](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deleterolealias)
  - [DeleteScheduledAudit](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletescheduledaudit)
  - [DeleteSecurityProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletesecurityprofile)
  - [DeleteStream](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletestream)
  - [DeleteThing](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletething)
  - [DeleteThingGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletethinggroup)
  - [DeleteThingType](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletethingtype)
  - [DeleteTopicRule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletetopicrule)
  - [DeleteTopicRuleDestination](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletetopicruledestination)
  - [DeleteV2LoggingLevel](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deletev2logginglevel)
  - [DeprecateThingType](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#deprecatethingtype)
  - [DescribeAccountAuditConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describeaccountauditconfiguration)
  - [DescribeAuditFinding](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describeauditfinding)
  - [DescribeAuditMitigationActionsTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describeauditmitigationactionstask)
  - [DescribeAuditSuppression](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describeauditsuppression)
  - [DescribeAuditTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describeaudittask)
  - [DescribeAuthorizer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describeauthorizer)
  - [DescribeBillingGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describebillinggroup)
  - [DescribeCACertificate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describecacertificate)
  - [DescribeCertificate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describecertificate)
  - [DescribeCertificateProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describecertificateprovider)
  - [DescribeCustomMetric](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describecustommetric)
  - [DescribeDefaultAuthorizer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describedefaultauthorizer)
  - [DescribeDetectMitigationActionsTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describedetectmitigationactionstask)
  - [DescribeDimension](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describedimension)
  - [DescribeDomainConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describedomainconfiguration)
  - [DescribeEncryptionConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describeencryptionconfiguration)
  - [DescribeEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describeendpoint)
  - [DescribeEventConfigurations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describeeventconfigurations)
  - [DescribeFleetMetric](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describefleetmetric)
  - [DescribeIndex](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describeindex)
  - [DescribeJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describejob)
  - [DescribeJobExecution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describejobexecution)
  - [DescribeJobTemplate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describejobtemplate)
  - [DescribeManagedJobTemplate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describemanagedjobtemplate)
  - [DescribeMitigationAction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describemitigationaction)
  - [DescribeProvisioningTemplate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describeprovisioningtemplate)
  - [DescribeProvisioningTemplateVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describeprovisioningtemplateversion)
  - [DescribeRoleAlias](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describerolealias)
  - [DescribeScheduledAudit](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describescheduledaudit)
  - [DescribeSecurityProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describesecurityprofile)
  - [DescribeStream](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describestream)
  - [DescribeThing](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describething)
  - [DescribeThingGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describethinggroup)
  - [DescribeThingRegistrationTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describethingregistrationtask)
  - [DescribeThingType](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#describethingtype)
  - [DetachPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#detachpolicy)
  - [DetachPrincipalPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#detachprincipalpolicy)
  - [DetachSecurityProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#detachsecurityprofile)
  - [DetachThingPrincipal](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#detachthingprincipal)
  - [DisableTopicRule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#disabletopicrule)
  - [DisassociateSbomFromPackageVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#disassociatesbomfrompackageversion)
  - [EnableTopicRule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#enabletopicrule)
  - [GetBehaviorModelTrainingSummaries](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#getbehaviormodeltrainingsummaries)
  - [GetBucketsAggregation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#getbucketsaggregation)
  - [GetCardinality](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#getcardinality)
  - [GetCommandExecution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#getcommandexecution)
  - [GetDeviceCommand](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#getdevicecommand)
  - [GetEffectivePolicies](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#geteffectivepolicies)
  - [GetIndexingConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#getindexingconfiguration)
  - [GetJobDocument](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#getjobdocument)
  - [GetLoggingOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#getloggingoptions)
  - [GetOTAUpdate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#getotaupdate)
  - [GetPackage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#getpackage)
  - [GetPackageConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#getpackageconfiguration)
  - [GetPackageVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#getpackageversion)
  - [GetPercentiles](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#getpercentiles)
  - [GetPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#getpolicy)
  - [GetPolicyVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#getpolicyversion)
  - [GetRegistrationCode](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#getregistrationcode)
  - [GetStatistics](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#getstatistics)
  - [GetThingConnectivityData](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#getthingconnectivitydata)
  - [GetTopicRule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#gettopicrule)
  - [GetTopicRuleDestination](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#gettopicruledestination)
  - [GetV2LoggingOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#getv2loggingoptions)
  - [ListActiveViolations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listactiveviolations)
  - [ListAttachedPolicies](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listattachedpolicies)
  - [ListAuditFindings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listauditfindings)
  - [ListAuditMitigationActionsExecutions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listauditmitigationactionsexecutions)
  - [ListAuditMitigationActionsTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listauditmitigationactionstasks)
  - [ListAuditSuppressions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listauditsuppressions)
  - [ListAuditTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listaudittasks)
  - [ListAuthorizers](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listauthorizers)
  - [ListBillingGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listbillinggroups)
  - [ListCACertificates](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listcacertificates)
  - [ListCertificateProviders](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listcertificateproviders)
  - [ListCertificates](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listcertificates)
  - [ListCertificatesByCA](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listcertificatesbyca)
  - [ListCommandExecutions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listcommandexecutions)
  - [ListCommands](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listcommands)
  - [ListCustomMetrics](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listcustommetrics)
  - [ListDetectMitigationActionsExecutions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listdetectmitigationactionsexecutions)
  - [ListDetectMitigationActionsTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listdetectmitigationactionstasks)
  - [ListDimensions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listdimensions)
  - [ListDomainConfigurations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listdomainconfigurations)
  - [ListFleetMetrics](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listfleetmetrics)
  - [ListIndices](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listindices)
  - [ListJobExecutionsForJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listjobexecutionsforjob)
  - [ListJobExecutionsForThing](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listjobexecutionsforthing)
  - [ListJobTemplates](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listjobtemplates)
  - [ListJobs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listjobs)
  - [ListManagedJobTemplates](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listmanagedjobtemplates)
  - [ListMetricValues](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listmetricvalues)
  - [ListMitigationActions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listmitigationactions)
  - [ListOTAUpdates](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listotaupdates)
  - [ListOutgoingCertificates](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listoutgoingcertificates)
  - [ListPackageVersions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listpackageversions)
  - [ListPackages](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listpackages)
  - [ListPolicies](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listpolicies)
  - [ListPolicyPrincipals](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listpolicyprincipals)
  - [ListPolicyVersions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listpolicyversions)
  - [ListPrincipalPolicies](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listprincipalpolicies)
  - [ListPrincipalThings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listprincipalthings)
  - [ListPrincipalThingsV2](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listprincipalthingsv2)
  - [ListProvisioningTemplateVersions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listprovisioningtemplateversions)
  - [ListProvisioningTemplates](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listprovisioningtemplates)
  - [ListRelatedResourcesForAuditFinding](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listrelatedresourcesforauditfinding)
  - [ListRoleAliases](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listrolealiases)
  - [ListSbomValidationResults](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listsbomvalidationresults)
  - [ListScheduledAudits](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listscheduledaudits)
  - [ListSecurityProfiles](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listsecurityprofiles)
  - [ListSecurityProfilesForTarget](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listsecurityprofilesfortarget)
  - [ListStreams](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#liststreams)
  - [ListTagsForResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listtagsforresource)
  - [ListTargetsForPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listtargetsforpolicy)
  - [ListTargetsForSecurityProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listtargetsforsecurityprofile)
  - [ListThingGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listthinggroups)
  - [ListThingGroupsForThing](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listthinggroupsforthing)
  - [ListThingPrincipals](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listthingprincipals)
  - [ListThingPrincipalsV2](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listthingprincipalsv2)
  - [ListThingRegistrationTaskReports](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listthingregistrationtaskreports)
  - [ListThingRegistrationTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listthingregistrationtasks)
  - [ListThingTypes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listthingtypes)
  - [ListThings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listthings)
  - [ListThingsInBillingGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listthingsinbillinggroup)
  - [ListThingsInThingGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listthingsinthinggroup)
  - [ListTopicRuleDestinations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listtopicruledestinations)
  - [ListTopicRules](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listtopicrules)
  - [ListV2LoggingLevels](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listv2logginglevels)
  - [ListViolationEvents](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#listviolationevents)
  - [PutVerificationStateOnViolation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#putverificationstateonviolation)
  - [RegisterCACertificate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#registercacertificate)
  - [RegisterCertificate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#registercertificate)
  - [RegisterCertificateWithoutCA](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#registercertificatewithoutca)
  - [RegisterThing](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#registerthing)
  - [RejectCertificateTransfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#rejectcertificatetransfer)
  - [RemoveThingFromBillingGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#removethingfrombillinggroup)
  - [RemoveThingFromThingGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#removethingfromthinggroup)
  - [ReplaceTopicRule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#replacetopicrule)
  - [SearchIndex](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#searchindex)
  - [SetDefaultAuthorizer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#setdefaultauthorizer)
  - [SetDefaultPolicyVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#setdefaultpolicyversion)
  - [SetLoggingOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#setloggingoptions)
  - [SetV2LoggingLevel](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#setv2logginglevel)
  - [SetV2LoggingOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#setv2loggingoptions)
  - [StartAuditMitigationActionsTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#startauditmitigationactionstask)
  - [StartDetectMitigationActionsTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#startdetectmitigationactionstask)
  - [StartOnDemandAuditTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#startondemandaudittask)
  - [StartThingRegistrationTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#startthingregistrationtask)
  - [StopThingRegistrationTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#stopthingregistrationtask)
  - [TagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#tagresource)
  - [TestAuthorization](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#testauthorization)
  - [TestInvokeAuthorizer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#testinvokeauthorizer)
  - [TransferCertificate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#transfercertificate)
  - [UntagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#untagresource)
  - [UpdateAccountAuditConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updateaccountauditconfiguration)
  - [UpdateAuditSuppression](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updateauditsuppression)
  - [UpdateAuthorizer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updateauthorizer)
  - [UpdateBillingGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updatebillinggroup)
  - [UpdateCACertificate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updatecacertificate)
  - [UpdateCertificate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updatecertificate)
  - [UpdateCertificateProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updatecertificateprovider)
  - [UpdateCommand](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updatecommand)
  - [UpdateCustomMetric](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updatecustommetric)
  - [UpdateDimension](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updatedimension)
  - [UpdateDomainConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updatedomainconfiguration)
  - [UpdateDynamicThingGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updatedynamicthinggroup)
  - [UpdateEncryptionConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updateencryptionconfiguration)
  - [UpdateEventConfigurations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updateeventconfigurations)
  - [UpdateFleetMetric](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updatefleetmetric)
  - [UpdateIndexingConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updateindexingconfiguration)
  - [UpdateJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updatejob)
  - [UpdateMitigationAction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updatemitigationaction)
  - [UpdatePackage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updatepackage)
  - [UpdatePackageConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updatepackageconfiguration)
  - [UpdatePackageVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updatepackageversion)
  - [UpdateProvisioningTemplate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updateprovisioningtemplate)
  - [UpdateRoleAlias](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updaterolealias)
  - [UpdateScheduledAudit](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updatescheduledaudit)
  - [UpdateSecurityProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updatesecurityprofile)
  - [UpdateStream](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updatestream)
  - [UpdateThing](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updatething)
  - [UpdateThingGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updatethinggroup)
  - [UpdateThingGroupsForThing](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updatethinggroupsforthing)
  - [UpdateThingType](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updatethingtype)
  - [UpdateTopicRuleDestination](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#updatetopicruledestination)
  - [ValidateSecurityProfileBehaviors](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iot-2015-05-28.html#validatesecurityprofilebehaviors)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Iot.IotClient.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Iot.IotClient.html\#toc-methods)

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

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Iot.IotClient.html\#methods)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Iot.IotClient.html#toc-methods)
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

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Iot.IotClient.html#top)
