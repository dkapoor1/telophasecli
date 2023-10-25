package cdktemplates

type CdkJSONData struct {
}

var CdkJSONTmpl = `
{
	"app": "go mod download && go run cdkapp.go",
	"watch": {
	  "include": [
		"**"
	  ],
	  "exclude": [
		"README.md",
		"cdk*.json",
		"go.mod",
		"go.sum",
		"**/*test.go"
	  ]
	},
	"context": {
	  "@aws-cdk/aws-lambda:recognizeLayerVersion": true,
	  "@aws-cdk/core:checkSecretUsage": true,
	  "@aws-cdk/core:target-partitions": [
		"aws",
		"aws-cn"
	  ],
	  "@aws-cdk-containers/ecs-service-extensions:enableDefaultLogDriver": true,
	  "@aws-cdk/aws-ec2:uniqueImdsv2TemplateName": true,
	  "@aws-cdk/aws-ecs:arnFormatIncludesClusterName": true,
	  "@aws-cdk/aws-iam:minimizePolicies": true,
	  "@aws-cdk/core:validateSnapshotRemovalPolicy": true,
	  "@aws-cdk/aws-codepipeline:crossAccountKeyAliasStackSafeResourceName": true,
	  "@aws-cdk/aws-s3:createDefaultLoggingPolicy": true,
	  "@aws-cdk/aws-sns-subscriptions:restrictSqsDescryption": true,
	  "@aws-cdk/aws-apigateway:disableCloudWatchRole": true,
	  "@aws-cdk/core:enablePartitionLiterals": true,
	  "@aws-cdk/aws-events:eventsTargetQueueSameAccount": true,
	  "@aws-cdk/aws-iam:standardizedServicePrincipals": true,
	  "@aws-cdk/aws-ecs:disableExplicitDeploymentControllerForCircuitBreaker": true,
	  "@aws-cdk/aws-iam:importedRoleStackSafeDefaultPolicyName": true,
	  "@aws-cdk/aws-s3:serverAccessLogsUseBucketPolicy": true,
	  "@aws-cdk/aws-route53-patters:useCertificate": true,
	  "@aws-cdk/customresources:installLatestAwsSdkDefault": false,
	  "@aws-cdk/aws-rds:databaseProxyUniqueResourceName": true,
	  "@aws-cdk/aws-codedeploy:removeAlarmsFromDeploymentGroup": true,
	  "@aws-cdk/aws-apigateway:authorizerChangeDeploymentLogicalId": true,
	  "@aws-cdk/aws-ec2:launchTemplateDefaultUserData": true,
	  "@aws-cdk/aws-secretsmanager:useAttachedSecretResourcePolicyForSecretTargetAttachments": true,
	  "@aws-cdk/aws-redshift:columnId": true,
	  "@aws-cdk/aws-stepfunctions-tasks:enableEmrServicePolicyV2": true,
	  "@aws-cdk/aws-ec2:restrictDefaultSecurityGroup": true,
	  "@aws-cdk/aws-apigateway:requestValidatorUniqueId": true,
	  "@aws-cdk/aws-kms:aliasNameRef": true,
	  "@aws-cdk/aws-autoscaling:generateLaunchTemplateInsteadOfLaunchConfig": true,
	  "@aws-cdk/core:includePrefixInUniqueNameGeneration": true,
	  "@aws-cdk/aws-efs:denyAnonymousAccess": true,
	  "@aws-cdk/aws-opensearchservice:enableOpensearchMultiAzWithStandby": true,
	  "@aws-cdk/aws-lambda-nodejs:useLatestRuntimeVersion": true,
	  "@aws-cdk/aws-efs:mountTargetOrderInsensitiveLogicalId": true,
	  "@aws-cdk/aws-rds:auroraClusterChangeScopeOfInstanceParameterGroupWithEachParameters": true,
	  "@aws-cdk/aws-appsync:useArnForSourceApiAssociationIdentifier": true,
	  "@aws-cdk/aws-rds:preventRenderingDeprecatedCredentials": true
	}
  }

`