# sacloud/security-control-api-go

Go言語向けのさくらのクラウド セキュリティコントロール APIライブラリ

マニュアル: https://manual.sakura.ad.jp/cloud/controlpanel/eventlog/security-control.html

APIドキュメント/OpenAPIは未公開なため、公開され次第更新。

## 概要

sacloud/security-control-api-goはさくらのクラウド セキュリティコントロール APIをGo言語から利用するためのAPIライブラリです。

```
package main

import (
	"context"
	"fmt"

	"github.com/sacloud/saclient-go"
	securitycontrol "github.com/sacloud/security-control-api-go"
	v1 "github.com/sacloud/security-control-api-go/apis/v1"
)

func runEvalRules() {
	var theClient saclient.Client
	client, err := securitycontrol.NewClient(&theClient)
	if err != nil {
		panic(err)
    }
    ctx := context.Background()

    // Evaluation Rules APIs
	erOp := securitycontrol.NewEvaluationRulesOp(client)
    updated, err := erOp.Update(ctx, "server-no-public-ip", v1.EvaluationRuleInput{
		IsEnabled: true,
		Rule: v1.EvaluationRuleUnion{
			OneOf: v1.EvaluationRuleUnionSum{
				Type: v1.EvaluationRuleUnionSumType("server-no-public-ip"),
				ServerNoPublicIP: v1.ServerNoPublicIP{
					EvaluationRuleId: v1.ServerNoPublicIPEvaluationRuleId("server-no-public-ip"),
					Parameter: v1.NewOptEvaluationRuleParametersZonedEvaluationTarget(
						v1.EvaluationRuleParametersZonedEvaluationTarget{
                            ServicePrincipalId: v1.NewOptString("111111111111"),
							Zones:              []string{"is1b", "tk1b"},
						},
					),
				},
			},
		},
	})
    // Read/List

    // Automated Actions API
    aaOp := securitycontrol.NewAutomatedActionsOp(client)
    created, err := aaOp.Create(ctx, &v1.AutomatedActionInput{
		Name:        "example-action",
		Description: v1.NewOptString("example automated action"),
		IsEnabled:   true,
		Action: v1.ActionDefinition{
			OneOf: v1.ActionDefinitionSum{
				Type: v1.ActionDefinitionSumType("simpleNotification"),
				ActionDefinitionSimpleNotification: v1.ActionDefinitionSimpleNotification{
					ActionType: v1.ActionDefinitionSimpleNotificationActionType("simpleNotification"),
					ActionParameter: v1.SakuraSimpleNotification{
						ServicePrincipalId:  "111111111111",
						NotificationGroupId: "111111111112",
					},
				},
			},
		},
		ExecutionCondition: `event.evaluationResult.status == "REJECTED"`,
	})
    // List/Read/Update/Delete
```

[evaluation_rules_test.go](./evaluation_rules_test.go) / [automated_actions_test.go](./automated_actions_test.go) 等も参照。

TODO: Automated Actions向けの設定ヘルパーを用意

:warning:  v1.0に達するまでは互換性のない形で変更される可能性がありますのでご注意ください。

## License

`security-control-api-go` Copyright (C) 2026- The sacloud/security-control-api-go authors.
This project is published under [Apache 2.0 License](LICENSE).

