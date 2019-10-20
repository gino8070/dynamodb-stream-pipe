package pipe

import (
	"encoding/json"
	"io"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
	"github.com/pkg/errors"
)

type App struct {
	table   string
	command string
	args    []string
	d       *dynamodb.DynamoDB
	ds      *dynamodbstreams.DynamoDBStreams
}

func NewApp(endpoint, table, command, args string) (*App, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String(endpoint),
	}))
	a := &App{
		table:   table,
		command: command,
		args:    strings.Split(args, ","),
		d:       dynamodb.New(sess),
		ds:      dynamodbstreams.New(sess),
	}
	return a, nil
}

func (a *App) Run() error {
	log.Println("run dynamodb streams piper")
	dto, err := a.d.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(a.table),
	})
	if err != nil {
		return errors.Wrap(err, "failed describe table")
	}
	if *dto.Table.LatestStreamArn == "" {
		return errors.New("disable dynamodb streams")
	}
	dso, err := a.ds.DescribeStream(&dynamodbstreams.DescribeStreamInput{
		StreamArn: dto.Table.LatestStreamArn,
	})
	if err != nil {
		return errors.Wrap(err, "failed describe stream")
	}
	gsio, err := a.ds.GetShardIterator(&dynamodbstreams.GetShardIteratorInput{
		ShardId:           dso.StreamDescription.Shards[len(dso.StreamDescription.Shards)-1].ShardId,
		StreamArn:         dto.Table.LatestStreamArn,
		ShardIteratorType: aws.String(dynamodbstreams.ShardIteratorTypeTrimHorizon),
	})
	if err != nil {
		return errors.Wrap(err, "failed get shard iterator")
	}
	itr := gsio.ShardIterator
	for {
		log.Printf("iterator %s", *itr)
		gro, err := a.ds.GetRecords(&dynamodbstreams.GetRecordsInput{
			ShardIterator: itr,
		})
		if err != nil {
			return errors.Wrap(err, "failed get records")
		}
		itr = gro.NextShardIterator
		log.Printf("num records: %d", len(gro.Records))
		for _, r := range gro.Records {
			records := &DynamoDBEvent{}
			nr := NewRecord(r)
			nr.EventSourceArn = dto.Table.LatestStreamArn
			records.Records = append(records.Records, nr)
			rj, _ := json.MarshalIndent(records, "", "  ")
			log.Printf("record: \n%s", rj)
			cmd := exec.Command(a.command, a.args...)
			stdin, _ := cmd.StdinPipe()
			io.WriteString(stdin, string(rj))
			stdin.Close()
			out, err := cmd.CombinedOutput()
			if err != nil {
				return errors.Wrap(err, "failed cmd")
			}
			log.Printf("cmd results: %s", out)
			time.Sleep(5 * time.Second)
		}
		if *itr == "" {
			break
		}
		time.Sleep(5 * time.Second)
	}
	return nil
}
