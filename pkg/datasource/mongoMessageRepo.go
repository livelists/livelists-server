package datasource

import (
	"fmt"
	"github.com/livelists/livelist-server/contracts/wsMessages"
	"github.com/livelists/livelist-server/pkg/config"
	"github.com/livelists/livelist-server/pkg/datasource/mongoSchemes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AddMessageArgs struct {
	ChannelIdentifier string
	SenderIdentifier  *string
	Text              string
	Type              string
	SubType           string
	CustomData        *wsMessages.CustomData
}

func AddMessage(args AddMessageArgs) (mongoSchemes.Message, error) {
	var client = config.GetMongoClient()

	fmt.Println("sender identifier", args.SenderIdentifier)

	newMessage := mongoSchemes.NewMessage(mongoSchemes.NewMessageArgs{
		ChannelIdentifier: args.ChannelIdentifier,
		SenderIdentifier:  args.SenderIdentifier,
		Text:              args.Text,
		Type:              args.Type,
		SubType:           args.SubType,
		CustomData:        args.CustomData,
	})

	_, err := client.Database(
		config.MainDatabase).Collection(
		mongoSchemes.MessageCollection).InsertOne(ctx, newMessage)

	if err != nil {
		return mongoSchemes.Message{}, err
	}

	return newMessage, nil
}

type GetMessagesFromChannelArgs struct {
	ChannelIdentifier string
	Skip              int
	Limit             int
	StartFromDate     time.Time
}
type GetMessagesFromChannelRes struct {
	Messages              *[]mongoSchemes.MessageWithParticipant
	TotalCount            int64
	FirstMessageCreatedAt time.Time
	LastMessageCreatedAt  time.Time
}

func GetMessagesFromChannel(args GetMessagesFromChannelArgs) (GetMessagesFromChannelRes, error) {
	var client = config.GetMongoClient()

	messages, err := client.Database(
		config.MainDatabase).Collection(mongoSchemes.MessageCollection).Aggregate(ctx, bson.A{
		bson.D{
			{"$match",
				bson.D{
					{"channel", args.ChannelIdentifier},
					{"createdAt", bson.D{{"$gte", primitive.NewDateTimeFromTime(args.StartFromDate)}}},
				},
			},
		},
		bson.D{{"$sort", bson.D{
			{"createdAt", -1},
		}}},
		bson.D{{"$skip", args.Skip}},
		bson.D{{"$limit", args.Limit}},
		bson.D{{"$sort", bson.D{{"createdAt", 1}}}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "Participant"},
					{"localField", "participant"},
					{"foreignField", "identifier"},
					{"as", "participant"},
				},
			},
		},
		bson.D{
			{"$addFields",
				bson.D{
					{"id", "$_id"},
					{"participant",
						bson.D{
							{"$arrayElemAt",
								bson.A{
									"$participant",
									0,
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{"$project",
				bson.D{
					{"_id", 0},
					{"id", 1},
					{"text", 1},
					{"customData", 1},
					{"channelIdentifier", args.ChannelIdentifier},
					{"type", 1},
					{"subType", 1},
					{"participant.identifier", 1},
					{"participant.customData", 1},
					{"createdAt", 1},
				},
			},
		},
		bson.D{{"$sort", bson.D{{"createdAt", 1}}}},
	})

	var messagesDocuments []mongoSchemes.MessageWithParticipant

	err = messages.All(ctx, &messagesDocuments)

	if err != nil {
		return GetMessagesFromChannelRes{}, err
	}

	var totalCount, errCount = CountMessagesInChannel(CountMessagesInChannelArgs{
		ChannelIdentifier: args.ChannelIdentifier,
	})

	if errCount != nil {
		return GetMessagesFromChannelRes{}, errCount
	}

	var lastAndFirstDates, datesErr = GetLastAndFirstMessagesCreatedAt(CountMessagesInChannelArgs{
		ChannelIdentifier: args.ChannelIdentifier,
	})

	if datesErr != nil {
		return GetMessagesFromChannelRes{}, errCount
	}

	return GetMessagesFromChannelRes{
		Messages:              &messagesDocuments,
		TotalCount:            totalCount,
		FirstMessageCreatedAt: lastAndFirstDates.FirstMessageCreatedAt,
		LastMessageCreatedAt:  lastAndFirstDates.LastMessageCreatedAt,
	}, nil
}

type CountMessagesInChannelArgs struct {
	ChannelIdentifier string
}

func CountMessagesInChannel(args CountMessagesInChannelArgs) (int64, error) {
	var client = config.GetMongoClient()

	messagesCount, err := client.
		Database(config.MainDatabase).
		Collection(mongoSchemes.MessageCollection).
		CountDocuments(ctx, bson.D{{"channel", args.ChannelIdentifier}})

	if err != nil {
		fmt.Println("Count messages error")
		return 0, err
	}

	return messagesCount, nil
}

func GetLastAndFirstMessagesCreatedAt(args CountMessagesInChannelArgs) (mongoSchemes.LastAndFirstMessagesCreatedAt, error) {
	var client = config.GetMongoClient()

	dates, err := client.Database(
		config.MainDatabase).Collection(mongoSchemes.MessageCollection).Aggregate(ctx, bson.A{
		bson.D{{"$match", bson.D{{"channel", args.ChannelIdentifier}}}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "Message"},
					{"as", "firstMessageCreatedAt"},
					{"let", bson.D{{"channelIdentifier", "$channel"}}},
					{"pipeline",
						bson.A{
							bson.D{
								{"$match",
									bson.D{
										{"$expr",
											bson.D{
												{"$eq",
													bson.A{
														"$channel",
														"$$channelIdentifier",
													},
												},
											},
										},
									},
								},
							},
							bson.D{{"$sort", bson.D{{"createdAt", 1}}}},
							bson.D{{"$limit", 1}},
						},
					},
				},
			},
		},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "Message"},
					{"as", "lastMessageCreatedAt"},
					{"let", bson.D{{"channelIdentifier", "$channel"}}},
					{"pipeline",
						bson.A{
							bson.D{
								{"$match",
									bson.D{
										{"$expr",
											bson.D{
												{"$eq",
													bson.A{
														"$channel",
														"$$channelIdentifier",
													},
												},
											},
										},
									},
								},
							},
							bson.D{{"$sort", bson.D{{"createdAt", -1}}}},
							bson.D{{"$limit", 1}},
						},
					},
				},
			},
		},
		bson.D{
			{"$addFields",
				bson.D{
					{"firstMessageCreatedAt",
						bson.D{
							{"$arrayElemAt",
								bson.A{
									"$firstMessageCreatedAt",
									0,
								},
							},
						},
					},
					{"lastMessageCreatedAt",
						bson.D{
							{"$arrayElemAt",
								bson.A{
									"$lastMessageCreatedAt",
									0,
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{"$addFields",
				bson.D{
					{"firstMessageCreatedAt", "$firstMessageCreatedAt.createdAt"},
					{"lastMessageCreatedAt", "$lastMessageCreatedAt.createdAt"},
				},
			},
		},
		bson.D{
			{"$project",
				bson.D{
					{"firstMessageCreatedAt", 1},
					{"lastMessageCreatedAt", 1},
				},
			},
		},
		bson.D{{"$limit", 1}},
	})

	var datesDocuments []mongoSchemes.LastAndFirstMessagesCreatedAt

	err = dates.All(ctx, &datesDocuments)

	if err != nil {
		return mongoSchemes.LastAndFirstMessagesCreatedAt{}, err
	}

	return datesDocuments[0], err
}

type CountMessagesInChannelAfterDateArgs struct {
	ChannelIdentifier string
	Date              time.Time
}

func CountMessagesInChannelAfterDate(args CountMessagesInChannelAfterDateArgs) (int64, error) {
	var client = config.GetMongoClient()

	messagesCount, err := client.
		Database(config.MainDatabase).
		Collection(mongoSchemes.MessageCollection).
		CountDocuments(ctx, bson.D{{
			"channel", args.ChannelIdentifier,
		}, {
			"createdAt", bson.D{{"$gt", primitive.NewDateTimeFromTime(args.Date)}},
		}})

	if err != nil {
		fmt.Println("Count messages error")
		return 0, err
	}

	return messagesCount, nil
}

type GetMessageCreatedAtByOffsetArgs struct {
	ChannelIdentifier string
	StartDate         time.Time
	Offset            int
}

func GetMessageCreatedAtByOffset(args GetMessageCreatedAtByOffsetArgs) (time.Time, error) {
	var client = config.GetMongoClient()

	messages, err := client.Database(
		config.MainDatabase).Collection(mongoSchemes.MessageCollection).Aggregate(ctx, bson.A{
		bson.D{
			{"$match",
				bson.D{
					{"channel", args.ChannelIdentifier},
					{"createdAt", bson.D{{"$lte", primitive.NewDateTimeFromTime(args.StartDate)}}},
				},
			},
		},
		bson.D{{"$sort", bson.D{{"createdAt", 1}}}},
		bson.D{{"$limit", args.Offset}},
		bson.D{{"$sort", bson.D{{"createdAt", -1}}}},
		bson.D{{"$limit", 1}},
	})

	var messagesDocuments []mongoSchemes.Message

	err = messages.All(ctx, &messagesDocuments)

	if err != nil {
		return args.StartDate, err
	}

	if len(messagesDocuments) == 0 {
		return args.StartDate, nil
	}

	return messagesDocuments[0].CreatedAt, err
}
