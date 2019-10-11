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
	Keys map[string]*AttributeValue `type:"map"`

	// The item in the DynamoDB table as it appeared after it was modified.
	NewImage map[string]*AttributeValue `type:"map"`

	// The item in the DynamoDB table as it appeared before it was modified.
	OldImage map[string]*AttributeValue `type:"map"`

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
		Keys:                        NewAttributeValue(r.Keys),
		NewImage:                    NewAttributeValue(r.NewImage),
		OldImage:                    NewAttributeValue(r.OldImage),
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

// Represents the data for an attribute.
//
// Each attribute value is described as a name-value pair. The name is the data
// type, and the value is the data itself.
//
// For more information, see Data Types (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.NamingRulesDataTypes.html#HowItWorks.DataTypes)
// in the Amazon DynamoDB Developer Guide.
type AttributeValue struct {
	_ struct{} `type:"structure"`

	// An attribute of type Binary. For example:
	//
	// "B": "dGhpcyB0ZXh0IGlzIGJhc2U2NC1lbmNvZGVk"
	//
	// B is automatically base64 encoded/decoded by the SDK.
	B []byte `type:"blob" json:"omitempty"`

	// An attribute of type Boolean. For example:
	//
	// "BOOL": true
	BOOL *bool `type:"boolean" json:"omitempty"`

	// An attribute of type Binary Set. For example:
	//
	// "BS": ["U3Vubnk=", "UmFpbnk=", "U25vd3k="]
	BS [][]byte `type:"list" json:"omitempty"`

	// An attribute of type List. For example:
	//
	// "L": [ {"S": "Cookies"} , {"S": "Coffee"}, {"N", "3.14159"}]
	L []*dynamodb.AttributeValue `type:"list" json:"omitempty"`

	// An attribute of type Map. For example:
	//
	// "M": {"Name": {"S": "Joe"}, "Age": {"N": "35"}}
	M map[string]*dynamodb.AttributeValue `type:"map" json:"omitempty"`

	// An attribute of type Number. For example:
	//
	// "N": "123.45"
	//
	// Numbers are sent across the network to DynamoDB as strings, to maximize compatibility
	// across languages and libraries. However, DynamoDB treats them as number type
	// attributes for mathematical operations.
	N *string `type:"string" json:"omitempty"`

	// An attribute of type Number Set. For example:
	//
	// "NS": ["42.2", "-19", "7.5", "3.14"]
	//
	// Numbers are sent across the network to DynamoDB as strings, to maximize compatibility
	// across languages and libraries. However, DynamoDB treats them as number type
	// attributes for mathematical operations.
	NS []*string `type:"list" json:"omitempty"`

	// An attribute of type Null. For example:
	//
	// "NULL": true
	NULL *bool `type:"boolean" json:"omitempty"`

	// An attribute of type String. For example:
	//
	// "S": "Hello"
	//S *string `type:"string" json:"omitempty"`
	S *string `type:"string"`

	// An attribute of type String Set. For example:
	//
	// "SS": ["Giraffe", "Hippo" ,"Zebra"]
	SS []*string `type:"list" json:"omitempty"`
}

func NewAttributeValue(m map[string]*dynamodb.AttributeValue) map[string]*AttributeValue {
	r := make(map[string]*AttributeValue)
	for k, v := range m {
		r[k] = &AttributeValue{
			B:    v.B,
			BOOL: v.BOOL,
			BS:   v.BS,
			L:    v.L,
			M:    v.M,
			N:    v.N,
			NS:   v.NS,
			NULL: v.NULL,
			S:    v.S,
			//S:  aws.String("hoge"),
			SS: v.SS,
		}
	}
	return r
}

// String returns the string representation
func (s AttributeValue) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AttributeValue) GoString() string {
	return s.String()
}
