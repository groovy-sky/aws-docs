Menu

- [Aws](namespace-aws.md)
- [Connect](namespace-aws-connect.md)

## ConnectClient     extends [AwsClient](class-aws-awsclient.md)   in package    - [Aws](package-aws.md)

This client is used to interact with the **Amazon Connect Service** service.

## Supported API Versions

This class uses a _service description model_ that is associated at
runtime based on the `version` option given when constructing the
client. The `version` option will determine which API operations,
waiters, and paginators are available for a client. Creating a command or a
specific API operation can be done using magic methods (e.g.,
`$client->commandName(/** parameters */)`, or using the
`$client->getCommand` method of the client.

- [**2017-08-08**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html)

  - [ActivateEvaluationForm](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#activateevaluationform)
  - [AssociateAnalyticsDataSet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#associateanalyticsdataset)
  - [AssociateApprovedOrigin](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#associateapprovedorigin)
  - [AssociateBot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#associatebot)
  - [AssociateContactWithUser](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#associatecontactwithuser)
  - [AssociateDefaultVocabulary](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#associatedefaultvocabulary)
  - [AssociateEmailAddressAlias](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#associateemailaddressalias)
  - [AssociateFlow](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#associateflow)
  - [AssociateHoursOfOperations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#associatehoursofoperations)
  - [AssociateInstanceStorageConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#associateinstancestorageconfig)
  - [AssociateLambdaFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#associatelambdafunction)
  - [AssociateLexBot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#associatelexbot)
  - [AssociatePhoneNumberContactFlow](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#associatephonenumbercontactflow)
  - [AssociateQueueEmailAddresses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#associatequeueemailaddresses)
  - [AssociateQueueQuickConnects](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#associatequeuequickconnects)
  - [AssociateRoutingProfileQueues](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#associateroutingprofilequeues)
  - [AssociateSecurityKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#associatesecuritykey)
  - [AssociateSecurityProfiles](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#associatesecurityprofiles)
  - [AssociateTrafficDistributionGroupUser](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#associatetrafficdistributiongroupuser)
  - [AssociateUserProficiencies](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#associateuserproficiencies)
  - [AssociateWorkspace](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#associateworkspace)
  - [BatchAssociateAnalyticsDataSet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#batchassociateanalyticsdataset)
  - [BatchCreateDataTableValue](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#batchcreatedatatablevalue)
  - [BatchDeleteDataTableValue](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#batchdeletedatatablevalue)
  - [BatchDescribeDataTableValue](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#batchdescribedatatablevalue)
  - [BatchDisassociateAnalyticsDataSet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#batchdisassociateanalyticsdataset)
  - [BatchGetAttachedFileMetadata](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#batchgetattachedfilemetadata)
  - [BatchGetFlowAssociation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#batchgetflowassociation)
  - [BatchPutContact](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#batchputcontact)
  - [BatchUpdateDataTableValue](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#batchupdatedatatablevalue)
  - [ClaimPhoneNumber](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#claimphonenumber)
  - [CompleteAttachedFileUpload](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#completeattachedfileupload)
  - [CreateAgentStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createagentstatus)
  - [CreateContact](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createcontact)
  - [CreateContactFlow](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createcontactflow)
  - [CreateContactFlowModule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createcontactflowmodule)
  - [CreateContactFlowModuleAlias](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createcontactflowmodulealias)
  - [CreateContactFlowModuleVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createcontactflowmoduleversion)
  - [CreateContactFlowVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createcontactflowversion)
  - [CreateDataTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createdatatable)
  - [CreateDataTableAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createdatatableattribute)
  - [CreateEmailAddress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createemailaddress)
  - [CreateEvaluationForm](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createevaluationform)
  - [CreateHoursOfOperation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createhoursofoperation)
  - [CreateHoursOfOperationOverride](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createhoursofoperationoverride)
  - [CreateInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createinstance)
  - [CreateIntegrationAssociation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createintegrationassociation)
  - [CreateNotification](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createnotification)
  - [CreateParticipant](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createparticipant)
  - [CreatePersistentContactAssociation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createpersistentcontactassociation)
  - [CreatePredefinedAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createpredefinedattribute)
  - [CreatePrompt](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createprompt)
  - [CreatePushNotificationRegistration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createpushnotificationregistration)
  - [CreateQueue](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createqueue)
  - [CreateQuickConnect](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createquickconnect)
  - [CreateRoutingProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createroutingprofile)
  - [CreateRule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createrule)
  - [CreateSecurityProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createsecurityprofile)
  - [CreateTaskTemplate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createtasktemplate)
  - [CreateTestCase](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createtestcase)
  - [CreateTrafficDistributionGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createtrafficdistributiongroup)
  - [CreateUseCase](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createusecase)
  - [CreateUser](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createuser)
  - [CreateUserHierarchyGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createuserhierarchygroup)
  - [CreateView](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createview)
  - [CreateViewVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createviewversion)
  - [CreateVocabulary](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createvocabulary)
  - [CreateWorkspace](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createworkspace)
  - [CreateWorkspacePage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#createworkspacepage)
  - [DeactivateEvaluationForm](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deactivateevaluationform)
  - [DeleteAttachedFile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deleteattachedfile)
  - [DeleteContactEvaluation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deletecontactevaluation)
  - [DeleteContactFlow](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deletecontactflow)
  - [DeleteContactFlowModule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deletecontactflowmodule)
  - [DeleteContactFlowModuleAlias](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deletecontactflowmodulealias)
  - [DeleteContactFlowModuleVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deletecontactflowmoduleversion)
  - [DeleteContactFlowVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deletecontactflowversion)
  - [DeleteDataTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deletedatatable)
  - [DeleteDataTableAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deletedatatableattribute)
  - [DeleteEmailAddress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deleteemailaddress)
  - [DeleteEvaluationForm](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deleteevaluationform)
  - [DeleteHoursOfOperation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deletehoursofoperation)
  - [DeleteHoursOfOperationOverride](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deletehoursofoperationoverride)
  - [DeleteInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deleteinstance)
  - [DeleteIntegrationAssociation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deleteintegrationassociation)
  - [DeleteNotification](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deletenotification)
  - [DeletePredefinedAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deletepredefinedattribute)
  - [DeletePrompt](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deleteprompt)
  - [DeletePushNotificationRegistration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deletepushnotificationregistration)
  - [DeleteQueue](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deletequeue)
  - [DeleteQuickConnect](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deletequickconnect)
  - [DeleteRoutingProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deleteroutingprofile)
  - [DeleteRule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deleterule)
  - [DeleteSecurityProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deletesecurityprofile)
  - [DeleteTaskTemplate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deletetasktemplate)
  - [DeleteTestCase](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deletetestcase)
  - [DeleteTrafficDistributionGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deletetrafficdistributiongroup)
  - [DeleteUseCase](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deleteusecase)
  - [DeleteUser](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deleteuser)
  - [DeleteUserHierarchyGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deleteuserhierarchygroup)
  - [DeleteView](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deleteview)
  - [DeleteViewVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deleteviewversion)
  - [DeleteVocabulary](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deletevocabulary)
  - [DeleteWorkspace](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deleteworkspace)
  - [DeleteWorkspaceMedia](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deleteworkspacemedia)
  - [DeleteWorkspacePage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#deleteworkspacepage)
  - [DescribeAgentStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describeagentstatus)
  - [DescribeAuthenticationProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describeauthenticationprofile)
  - [DescribeContact](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describecontact)
  - [DescribeContactEvaluation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describecontactevaluation)
  - [DescribeContactFlow](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describecontactflow)
  - [DescribeContactFlowModule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describecontactflowmodule)
  - [DescribeContactFlowModuleAlias](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describecontactflowmodulealias)
  - [DescribeDataTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describedatatable)
  - [DescribeDataTableAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describedatatableattribute)
  - [DescribeEmailAddress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describeemailaddress)
  - [DescribeEvaluationForm](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describeevaluationform)
  - [DescribeHoursOfOperation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describehoursofoperation)
  - [DescribeHoursOfOperationOverride](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describehoursofoperationoverride)
  - [DescribeInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describeinstance)
  - [DescribeInstanceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describeinstanceattribute)
  - [DescribeInstanceStorageConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describeinstancestorageconfig)
  - [DescribeNotification](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describenotification)
  - [DescribePhoneNumber](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describephonenumber)
  - [DescribePredefinedAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describepredefinedattribute)
  - [DescribePrompt](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describeprompt)
  - [DescribeQueue](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describequeue)
  - [DescribeQuickConnect](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describequickconnect)
  - [DescribeRoutingProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describeroutingprofile)
  - [DescribeRule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describerule)
  - [DescribeSecurityProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describesecurityprofile)
  - [DescribeTestCase](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describetestcase)
  - [DescribeTrafficDistributionGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describetrafficdistributiongroup)
  - [DescribeUser](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describeuser)
  - [DescribeUserHierarchyGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describeuserhierarchygroup)
  - [DescribeUserHierarchyStructure](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describeuserhierarchystructure)
  - [DescribeView](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describeview)
  - [DescribeVocabulary](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describevocabulary)
  - [DescribeWorkspace](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#describeworkspace)
  - [DisassociateAnalyticsDataSet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#disassociateanalyticsdataset)
  - [DisassociateApprovedOrigin](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#disassociateapprovedorigin)
  - [DisassociateBot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#disassociatebot)
  - [DisassociateEmailAddressAlias](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#disassociateemailaddressalias)
  - [DisassociateFlow](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#disassociateflow)
  - [DisassociateHoursOfOperations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#disassociatehoursofoperations)
  - [DisassociateInstanceStorageConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#disassociateinstancestorageconfig)
  - [DisassociateLambdaFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#disassociatelambdafunction)
  - [DisassociateLexBot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#disassociatelexbot)
  - [DisassociatePhoneNumberContactFlow](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#disassociatephonenumbercontactflow)
  - [DisassociateQueueEmailAddresses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#disassociatequeueemailaddresses)
  - [DisassociateQueueQuickConnects](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#disassociatequeuequickconnects)
  - [DisassociateRoutingProfileQueues](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#disassociateroutingprofilequeues)
  - [DisassociateSecurityKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#disassociatesecuritykey)
  - [DisassociateSecurityProfiles](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#disassociatesecurityprofiles)
  - [DisassociateTrafficDistributionGroupUser](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#disassociatetrafficdistributiongroupuser)
  - [DisassociateUserProficiencies](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#disassociateuserproficiencies)
  - [DisassociateWorkspace](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#disassociateworkspace)
  - [DismissUserContact](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#dismissusercontact)
  - [EvaluateDataTableValues](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#evaluatedatatablevalues)
  - [GetAttachedFile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#getattachedfile)
  - [GetContactAttributes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#getcontactattributes)
  - [GetContactMetrics](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#getcontactmetrics)
  - [GetCurrentMetricData](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#getcurrentmetricdata)
  - [GetCurrentUserData](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#getcurrentuserdata)
  - [GetEffectiveHoursOfOperations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#geteffectivehoursofoperations)
  - [GetFederationToken](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#getfederationtoken)
  - [GetFlowAssociation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#getflowassociation)
  - [GetMetricData](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#getmetricdata)
  - [GetMetricDataV2](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#getmetricdatav2)
  - [GetPromptFile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#getpromptfile)
  - [GetTaskTemplate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#gettasktemplate)
  - [GetTestCaseExecutionSummary](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#gettestcaseexecutionsummary)
  - [GetTrafficDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#gettrafficdistribution)
  - [ImportPhoneNumber](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#importphonenumber)
  - [ImportWorkspaceMedia](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#importworkspacemedia)
  - [ListAgentStatuses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listagentstatuses)
  - [ListAnalyticsDataAssociations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listanalyticsdataassociations)
  - [ListAnalyticsDataLakeDataSets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listanalyticsdatalakedatasets)
  - [ListApprovedOrigins](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listapprovedorigins)
  - [ListAssociatedContacts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listassociatedcontacts)
  - [ListAuthenticationProfiles](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listauthenticationprofiles)
  - [ListBots](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listbots)
  - [ListChildHoursOfOperations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listchildhoursofoperations)
  - [ListContactEvaluations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listcontactevaluations)
  - [ListContactFlowModuleAliases](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listcontactflowmodulealiases)
  - [ListContactFlowModuleVersions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listcontactflowmoduleversions)
  - [ListContactFlowModules](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listcontactflowmodules)
  - [ListContactFlowVersions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listcontactflowversions)
  - [ListContactFlows](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listcontactflows)
  - [ListContactReferences](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listcontactreferences)
  - [ListDataTableAttributes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listdatatableattributes)
  - [ListDataTablePrimaryValues](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listdatatableprimaryvalues)
  - [ListDataTableValues](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listdatatablevalues)
  - [ListDataTables](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listdatatables)
  - [ListDefaultVocabularies](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listdefaultvocabularies)
  - [ListEntitySecurityProfiles](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listentitysecurityprofiles)
  - [ListEvaluationFormVersions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listevaluationformversions)
  - [ListEvaluationForms](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listevaluationforms)
  - [ListFlowAssociations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listflowassociations)
  - [ListHoursOfOperationOverrides](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listhoursofoperationoverrides)
  - [ListHoursOfOperations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listhoursofoperations)
  - [ListInstanceAttributes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listinstanceattributes)
  - [ListInstanceStorageConfigs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listinstancestorageconfigs)
  - [ListInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listinstances)
  - [ListIntegrationAssociations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listintegrationassociations)
  - [ListLambdaFunctions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listlambdafunctions)
  - [ListLexBots](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listlexbots)
  - [ListNotifications](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listnotifications)
  - [ListPhoneNumbers](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listphonenumbers)
  - [ListPhoneNumbersV2](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listphonenumbersv2)
  - [ListPredefinedAttributes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listpredefinedattributes)
  - [ListPrompts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listprompts)
  - [ListQueueEmailAddresses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listqueueemailaddresses)
  - [ListQueueQuickConnects](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listqueuequickconnects)
  - [ListQueues](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listqueues)
  - [ListQuickConnects](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listquickconnects)
  - [ListRealtimeContactAnalysisSegmentsV2](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listrealtimecontactanalysissegmentsv2)
  - [ListRoutingProfileManualAssignmentQueues](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listroutingprofilemanualassignmentqueues)
  - [ListRoutingProfileQueues](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listroutingprofilequeues)
  - [ListRoutingProfiles](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listroutingprofiles)
  - [ListRules](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listrules)
  - [ListSecurityKeys](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listsecuritykeys)
  - [ListSecurityProfileApplications](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listsecurityprofileapplications)
  - [ListSecurityProfileFlowModules](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listsecurityprofileflowmodules)
  - [ListSecurityProfilePermissions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listsecurityprofilepermissions)
  - [ListSecurityProfiles](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listsecurityprofiles)
  - [ListTagsForResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listtagsforresource)
  - [ListTaskTemplates](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listtasktemplates)
  - [ListTestCaseExecutionRecords](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listtestcaseexecutionrecords)
  - [ListTestCaseExecutions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listtestcaseexecutions)
  - [ListTestCases](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listtestcases)
  - [ListTrafficDistributionGroupUsers](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listtrafficdistributiongroupusers)
  - [ListTrafficDistributionGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listtrafficdistributiongroups)
  - [ListUseCases](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listusecases)
  - [ListUserHierarchyGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listuserhierarchygroups)
  - [ListUserNotifications](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listusernotifications)
  - [ListUserProficiencies](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listuserproficiencies)
  - [ListUsers](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listusers)
  - [ListViewVersions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listviewversions)
  - [ListViews](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listviews)
  - [ListWorkspaceMedia](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listworkspacemedia)
  - [ListWorkspacePages](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listworkspacepages)
  - [ListWorkspaces](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#listworkspaces)
  - [MonitorContact](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#monitorcontact)
  - [PauseContact](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#pausecontact)
  - [PutUserStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#putuserstatus)
  - [ReleasePhoneNumber](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#releasephonenumber)
  - [ReplicateInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#replicateinstance)
  - [ResumeContact](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#resumecontact)
  - [ResumeContactRecording](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#resumecontactrecording)
  - [SearchAgentStatuses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchagentstatuses)
  - [SearchAvailablePhoneNumbers](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchavailablephonenumbers)
  - [SearchContactEvaluations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchcontactevaluations)
  - [SearchContactFlowModules](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchcontactflowmodules)
  - [SearchContactFlows](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchcontactflows)
  - [SearchContacts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchcontacts)
  - [SearchDataTables](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchdatatables)
  - [SearchEmailAddresses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchemailaddresses)
  - [SearchEvaluationForms](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchevaluationforms)
  - [SearchHoursOfOperationOverrides](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchhoursofoperationoverrides)
  - [SearchHoursOfOperations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchhoursofoperations)
  - [SearchNotifications](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchnotifications)
  - [SearchPredefinedAttributes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchpredefinedattributes)
  - [SearchPrompts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchprompts)
  - [SearchQueues](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchqueues)
  - [SearchQuickConnects](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchquickconnects)
  - [SearchResourceTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchresourcetags)
  - [SearchRoutingProfiles](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchroutingprofiles)
  - [SearchSecurityProfiles](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchsecurityprofiles)
  - [SearchTestCases](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchtestcases)
  - [SearchUserHierarchyGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchuserhierarchygroups)
  - [SearchUsers](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchusers)
  - [SearchViews](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchviews)
  - [SearchVocabularies](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchvocabularies)
  - [SearchWorkspaceAssociations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchworkspaceassociations)
  - [SearchWorkspaces](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#searchworkspaces)
  - [SendChatIntegrationEvent](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#sendchatintegrationevent)
  - [SendOutboundEmail](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#sendoutboundemail)
  - [StartAttachedFileUpload](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#startattachedfileupload)
  - [StartChatContact](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#startchatcontact)
  - [StartContactEvaluation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#startcontactevaluation)
  - [StartContactMediaProcessing](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#startcontactmediaprocessing)
  - [StartContactRecording](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#startcontactrecording)
  - [StartContactStreaming](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#startcontactstreaming)
  - [StartEmailContact](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#startemailcontact)
  - [StartOutboundChatContact](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#startoutboundchatcontact)
  - [StartOutboundEmailContact](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#startoutboundemailcontact)
  - [StartOutboundVoiceContact](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#startoutboundvoicecontact)
  - [StartScreenSharing](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#startscreensharing)
  - [StartTaskContact](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#starttaskcontact)
  - [StartTestCaseExecution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#starttestcaseexecution)
  - [StartWebRTCContact](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#startwebrtccontact)
  - [StopContact](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#stopcontact)
  - [StopContactMediaProcessing](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#stopcontactmediaprocessing)
  - [StopContactRecording](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#stopcontactrecording)
  - [StopContactStreaming](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#stopcontactstreaming)
  - [StopTestCaseExecution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#stoptestcaseexecution)
  - [SubmitContactEvaluation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#submitcontactevaluation)
  - [SuspendContactRecording](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#suspendcontactrecording)
  - [TagContact](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#tagcontact)
  - [TagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#tagresource)
  - [TransferContact](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#transfercontact)
  - [UntagContact](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#untagcontact)
  - [UntagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#untagresource)
  - [UpdateAgentStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateagentstatus)
  - [UpdateAuthenticationProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateauthenticationprofile)
  - [UpdateContact](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatecontact)
  - [UpdateContactAttributes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatecontactattributes)
  - [UpdateContactEvaluation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatecontactevaluation)
  - [UpdateContactFlowContent](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatecontactflowcontent)
  - [UpdateContactFlowMetadata](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatecontactflowmetadata)
  - [UpdateContactFlowModuleAlias](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatecontactflowmodulealias)
  - [UpdateContactFlowModuleContent](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatecontactflowmodulecontent)
  - [UpdateContactFlowModuleMetadata](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatecontactflowmodulemetadata)
  - [UpdateContactFlowName](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatecontactflowname)
  - [UpdateContactRoutingData](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatecontactroutingdata)
  - [UpdateContactSchedule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatecontactschedule)
  - [UpdateDataTableAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatedatatableattribute)
  - [UpdateDataTableMetadata](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatedatatablemetadata)
  - [UpdateDataTablePrimaryValues](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatedatatableprimaryvalues)
  - [UpdateEmailAddressMetadata](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateemailaddressmetadata)
  - [UpdateEvaluationForm](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateevaluationform)
  - [UpdateHoursOfOperation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatehoursofoperation)
  - [UpdateHoursOfOperationOverride](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatehoursofoperationoverride)
  - [UpdateInstanceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateinstanceattribute)
  - [UpdateInstanceStorageConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateinstancestorageconfig)
  - [UpdateNotificationContent](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatenotificationcontent)
  - [UpdateParticipantAuthentication](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateparticipantauthentication)
  - [UpdateParticipantRoleConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateparticipantroleconfig)
  - [UpdatePhoneNumber](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatephonenumber)
  - [UpdatePhoneNumberMetadata](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatephonenumbermetadata)
  - [UpdatePredefinedAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatepredefinedattribute)
  - [UpdatePrompt](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateprompt)
  - [UpdateQueueHoursOfOperation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatequeuehoursofoperation)
  - [UpdateQueueMaxContacts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatequeuemaxcontacts)
  - [UpdateQueueName](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatequeuename)
  - [UpdateQueueOutboundCallerConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatequeueoutboundcallerconfig)
  - [UpdateQueueOutboundEmailConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatequeueoutboundemailconfig)
  - [UpdateQueueStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatequeuestatus)
  - [UpdateQuickConnectConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatequickconnectconfig)
  - [UpdateQuickConnectName](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatequickconnectname)
  - [UpdateRoutingProfileAgentAvailabilityTimer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateroutingprofileagentavailabilitytimer)
  - [UpdateRoutingProfileConcurrency](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateroutingprofileconcurrency)
  - [UpdateRoutingProfileDefaultOutboundQueue](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateroutingprofiledefaultoutboundqueue)
  - [UpdateRoutingProfileName](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateroutingprofilename)
  - [UpdateRoutingProfileQueues](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateroutingprofilequeues)
  - [UpdateRule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updaterule)
  - [UpdateSecurityProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatesecurityprofile)
  - [UpdateTaskTemplate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatetasktemplate)
  - [UpdateTestCase](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatetestcase)
  - [UpdateTrafficDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updatetrafficdistribution)
  - [UpdateUserConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateuserconfig)
  - [UpdateUserHierarchy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateuserhierarchy)
  - [UpdateUserHierarchyGroupName](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateuserhierarchygroupname)
  - [UpdateUserHierarchyStructure](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateuserhierarchystructure)
  - [UpdateUserIdentityInfo](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateuseridentityinfo)
  - [UpdateUserNotificationStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateusernotificationstatus)
  - [UpdateUserPhoneConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateuserphoneconfig)
  - [UpdateUserProficiencies](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateuserproficiencies)
  - [UpdateUserRoutingProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateuserroutingprofile)
  - [UpdateUserSecurityProfiles](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateusersecurityprofiles)
  - [UpdateViewContent](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateviewcontent)
  - [UpdateViewMetadata](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateviewmetadata)
  - [UpdateWorkspaceMetadata](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateworkspacemetadata)
  - [UpdateWorkspacePage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateworkspacepage)
  - [UpdateWorkspaceTheme](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateworkspacetheme)
  - [UpdateWorkspaceVisibility](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-connect-2017-08-08.html#updateworkspacevisibility)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Connect.ConnectClient.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Connect.ConnectClient.html\#toc-methods)

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

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Connect.ConnectClient.html\#methods)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Connect.ConnectClient.html#toc-methods)
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

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Connect.ConnectClient.html#top)
