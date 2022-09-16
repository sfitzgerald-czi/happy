package aws

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	ecstypes "github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/chanzuckerberg/happy/pkg/backend/aws/interfaces"
	"github.com/chanzuckerberg/happy/pkg/config"
	"github.com/chanzuckerberg/happy/pkg/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestNetworkConfig(t *testing.T) {
	r := require.New(t)
	backend := Backend{}
	sgs := []string{"sg-1", "sg-2"}
	subnets := []string{"subnet-1", "subnet-2"}

	backend.integrationSecret = &config.IntegrationSecret{
		ClusterArn:     "arn:cluster",
		PrivateSubnets: subnets,
		SecurityGroups: sgs,
		Services:       map[string]*config.RegistryConfig{},
	}
	networkConfig := backend.getNetworkConfig()
	r.NotNil(networkConfig)
	r.Equal(len(subnets), len(networkConfig.AwsvpcConfiguration.Subnets))
	r.Equal(len(sgs), len(networkConfig.AwsvpcConfiguration.SecurityGroups))

	for index, subnet := range subnets {
		r.Equal(subnet, networkConfig.AwsvpcConfiguration.Subnets[index])
	}
	for index, sg := range sgs {
		r.Equal(sg, networkConfig.AwsvpcConfiguration.SecurityGroups[index])
	}
}

func TestWaitForTasks(t *testing.T) {
	r := require.New(t)

	ctx := context.WithValue(context.Background(), util.CmdStartContextKey, time.Now())

	ctrl := gomock.NewController(t)

	bootstrapConfig := &config.Bootstrap{
		HappyConfigPath:         testFilePath,
		DockerComposeConfigPath: testDockerComposePath,
		Env:                     "rdev",
	}

	happyConfig, err := config.NewHappyConfig(bootstrapConfig)
	r.NoError(err)

	secretsApi := interfaces.NewMockSecretsManagerAPI(ctrl)
	testVal := "{\"cluster_arn\": \"test_arn\",\"ecrs\": {\"ecr_1\": {\"url\": \"test_url_1\"}},\"tfe\": {\"url\": \"tfe_url\",\"org\": \"tfe_org\"}}"
	secretsApi.EXPECT().GetSecretValue(gomock.Any(), gomock.Any()).
		Return(&secretsmanager.GetSecretValueOutput{
			SecretString: &testVal,
			ARN:          aws.String("arn:aws:secretsmanager:region:accountid:secret:happy/env-happy-config-AB1234"),
		}, nil).AnyTimes()

	stsApi := interfaces.NewMockSTSAPI(ctrl)
	stsApi.EXPECT().GetCallerIdentity(gomock.Any(), gomock.Any()).
		Return(&sts.GetCallerIdentityOutput{UserId: aws.String("foo:bar")}, nil).AnyTimes()

	tasks := []ecstypes.Task{}
	startedAt := time.Now().Add(time.Duration(-2) * time.Hour)

	containers := []ecstypes.Container{}
	containers = append(containers, ecstypes.Container{
		Name:      aws.String("nginx"),
		RuntimeId: aws.String("123"),
		TaskArn:   aws.String("arn:::::ecs/task/name/mytaskid"),
	})

	tasks = append(tasks, ecstypes.Task{
		LastStatus:           aws.String("RUNNING"),
		ContainerInstanceArn: aws.String("host"),
		StartedAt:            &startedAt,
		Containers:           containers,
		LaunchType:           ecstypes.LaunchTypeEc2,
		TaskDefinitionArn:    aws.String("arn:aws:ecs:us-west-2:123456789012:task-definition/hello_world:8"),
		TaskArn:              aws.String("arn:::::ecs/task/name/mytaskid"),
	})

	ecsApi := interfaces.NewMockECSAPI(ctrl)
	ecsApi.EXPECT().DescribeTasks(gomock.Any(), gomock.Any()).Return(&ecs.DescribeTasksOutput{
		Failures: []ecstypes.Failure{},
		Tasks:    tasks,
	}, nil).AnyTimes()

	taskStoppedWaiter := interfaces.NewMockECSTaskStoppedWaiterAPI(ctrl)
	taskStoppedWaiter.EXPECT().Wait(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	b, err := NewAWSBackend(ctx, happyConfig,
		WithAWSAccountID("1234567890"),
		WithSTSClient(stsApi),
		WithSecretsClient(secretsApi),
		WithECSClient(ecsApi),
		WithTaskStoppedWaiter(taskStoppedWaiter),
	)
	r.NoError(err)
	_, err = b.waitForTasksToStart(ctx, []string{"arn:::::ecs/task/name/mytaskid"})
	r.NoError(err)
	err = b.waitForTasksToStop(ctx, []string{"arn:::::ecs/task/name/mytaskid"})
	r.NoError(err)
}
