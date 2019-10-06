package pipe

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
)

// A description of a unique event within a stream.
type Record struct {
	_ struct{} `type:"structure"`

	// The region in which the GetRecords request was received.
	AwsRegion *string `locationName:"awsRegion" type:"string"`

	// The main body of the stream record, containing all of the DynamoDB-specific
	// fields.
	Dynamodb *StreamRecord `locationName:"dynamodb" type:"structure"`

	// A globally unique identifier for the event that was recorded in this stream
	// record.
	EventID *string `locationName:"eventID" type:"string"`

	// The type of data modification that was performed on the DynamoDB table:
	//
	//    * INSERT - a new item was added to the table.
	//
	//    * MODIFY - one or more of an existing item's attributes were modified.
	//
	//    * REMOVE - the item was deleted from the table
	EventName *string `locationName:"eventName" type:"string" enum:"OperationType"`

	// The AWS service from which the stream record originated. For DynamoDB Streams,
	// this is aws:dynamodb.
	EventSource *string `locationName:"eventSource" type:"string"`

	// The version number of the stream record format. This number is updated whenever
	// the structure of Record is modified.
	//
	// Client applications must not assume that eventVersion will remain at a particular
	// value, as this number is subject to change at any time. In general, eventVersion
	// will only increase as the low-level DynamoDB Streams API evolves.
	EventVersion *string `locationName:"eventVersion" type:"string"`

	// Items that are deleted by the Time to Live process after expiration have
	// the following fields:
	//
	//    * Records[].userIdentity.type "Service"
	//
	//    * Records[].userIdentity.principalId "dynamodb.amazonaws.com"
	UserIdentity *dynamodbstreams.Identity `locationName:"userIdentity" type:"structure"`
}

// String returns the string representation
func (s Record) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Record) GoString() string {
	return s.String()
}

func NewRecord(r *dynamodbstreams.Record) *Record {
	return &Record{
		AwsRegion:    r.AwsRegion,
		Dynamodb:     NewStreamRecord(r.Dynamodb),
		EventID:      r.EventID,
		EventName:    r.EventName,
		EventSource:  r.EventSource,
		EventVersion: r.EventVersion,
		UserIdentity: r.UserIdentity,
	}
}

// A description of a single data modification that was performed on an item
// in a DynamoDB table.
type StreamRecord struct {
	_ struct{} `type:"structure"`

	// The approximate date and time when the stream record was created, in UNIX
	// epoch time (http://www.epochconverter.com/) format.
	ApproximateCreationDateTime *int64 `type:"long"`

	// The primary key attribute(s) for the DynamoDB item that was modified.
	Keys map[string]*dynamodb.AttributeValue `type:"map"`

	// The item in the DynamoDB table as it appeared after it was modified.
	NewImage map[string]*dynamodb.AttributeValue `type:"map"`

	// The item in the DynamoDB table as it appeared before it was modified.
	OldImage map[string]*dynamodb.AttributeValue `type:"map"`

	// The sequence number of the stream record.
	SequenceNumber *string `min:"21" type:"string"`

	// The size of the stream record, in bytes.
	SizeBytes *int64 `min:"1" type:"long"`

	// The type of data from the modified DynamoDB item that was captured in this
	// stream record:
	//
	//    * KEYS_ONLY - only the key attributes of the modified item.
	//
	//    * NEW_IMAGE - the entire item, as it appeared after it was modified.
	//
	//    * OLD_IMAGE - the entire item, as it appeared before it was modified.
	//
	//    * NEW_AND_OLD_IMAGES - both the new and the old item images of the item.
	StreamViewType *string `type:"string" enum:"StreamViewType"`
}

func NewStreamRecord(r *dynamodbstreams.StreamRecord) *StreamRecord {
	return &StreamRecord{
		ApproximateCreationDateTime: aws.Int64(r.ApproximateCreationDateTime.Unix()),
		Keys:                        r.Keys,
		NewImage:                    r.NewImage,
		OldImage:                    r.OldImage,
		SequenceNumber:              r.SequenceNumber,
		SizeBytes:                   r.SizeBytes,
	}
}

// String returns the string representation
func (s StreamRecord) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StreamRecord) GoString() string {
	return s.String()
}
