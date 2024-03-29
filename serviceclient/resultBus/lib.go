package resultBus

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"

	"github.com/provenian/api/functions/submit/model"
)

type ResultBus struct {
	topicArn string
	*sns.SNS
}

func NewResultBus(topicArn string, instance *sns.SNS) ResultBus {
	return ResultBus{
		topicArn: topicArn,
		SNS:      instance,
	}
}

func (bus *ResultBus) Send(message ResultMessage) error {
	out, err := json.Marshal(&message)
	if err != nil {
		return err
	}

	if _, err := bus.SNS.Publish(&sns.PublishInput{
		TopicArn: aws.String(bus.topicArn),
		Message:  aws.String(string(out)),
	}); err != nil {
		return err
	}

	return nil
}

type ResultMessage struct {
	SubmissionID string       `json:"submission_id"`
	ProblemID    string       `json:"problem_id"`
	UserID       string       `json:"user_id"`
	Language     string       `json:"language"`
	Result       model.Result `json:"result"`
}

func NewResultMessage(submissionID string, problemID string, userID string, result model.Result) ResultMessage {
	return ResultMessage{
		SubmissionID: submissionID,
		ProblemID:    problemID,
		UserID:       userID,
		Result:       result,
	}
}
