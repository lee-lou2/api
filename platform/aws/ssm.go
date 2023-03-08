package aws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"log"
	"os"
)

func SetEnvironments() {
	cfg, err := awsConfig()
	if err != nil {
		panic(fmt.Sprintf("환경 변수 설정을 위한 AWS 연결이 실패하였습니다. 오류 내용 : %s\n", err.Error()))
	}

	// SSM 서비스 클라이언트 생성
	svc := ssm.NewFromConfig(cfg)

	// SSM 파라미터 가져오기
	var paramsGroup [][]types.Parameter
	var nextToken *string
	for {
		input := &ssm.GetParametersByPathInput{
			Path:           aws.String("/api/prod"),
			Recursive:      aws.Bool(true),
			WithDecryption: aws.Bool(true),
			NextToken:      nextToken,
		}
		resp, err := svc.GetParametersByPath(context.Background(), input)
		if err != nil {
			panic("failed to get SSM parameters")
		}
		paramsGroup = append(paramsGroup, resp.Parameters)
		if resp.NextToken == nil {
			break
		}
		nextToken = resp.NextToken
	}

	// 환경 변수에 파라미터 값 할당
	paramsCnt := 0
	for _, params := range paramsGroup {
		for _, param := range params {
			paramName := *param.Name
			paramValue := *param.Value
			envName := paramName[len("/api/prod/"):]
			os.Setenv(envName, paramValue)
			paramsCnt += 1
		}
	}
	log.Printf("%d 개의 환경 변수 설정 완료\n", paramsCnt)
}
