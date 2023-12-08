package datasource

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "github.com/livelists/livelist-server/contracts/channel"
	"github.com/livelists/livelist-server/pkg/config"
	"github.com/livelists/livelist-server/pkg/datasource/mongoSchemes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CreateChannelArgs struct {
	Identifier      string
	MaxParticipants int64
}

func CreateChannel(args CreateChannelArgs) pb.Channel {
	var client = config.GetMongoClient()

	newChannel := mongoSchemes.NewChannel(mongoSchemes.NewChannelArgs{
		Identifier:      args.Identifier,
		Status:          pb.ChannelStatus_Active,
		MaxParticipants: args.MaxParticipants,
	})

	_, err := client.Database(config.MainDatabase).Collection(mongoSchemes.ChannelCollection).InsertOne(ctx, newChannel)

	fmt.Print("channel create", err)

	return pb.Channel{
		Identifier:      args.Identifier,
		Status:          pb.ChannelStatus_Active,
		MaxParticipants: args.MaxParticipants,
		CreatedAt:       &timestamp.Timestamp{Seconds: int64(time.Now().Second())},
	}
}

type GetParticipantChWithMsgArgs struct {
	MessagesLimit         int32
	ParticipantIdentifier string
}

func GetParticipantsChannelsWithMessages(args GetParticipantChWithMsgArgs) ([]mongoSchemes.ChannelWithLastMessages, error) {
	var client = config.GetMongoClient()

	channelsWithMsg, err := client.Database(
		config.MainDatabase).Collection(mongoSchemes.ParticipantCollection).Aggregate(ctx, bson.A{
		bson.D{{"$match", bson.D{{
			"identifier", args.ParticipantIdentifier,
		}}}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "Channel"},
					{"localField", "channel"},
					{"foreignField", "identifier"},
					{"as", "channel"},
				},
			},
		},
		bson.D{
			{"$addFields",
				bson.D{
					{"channel",
						bson.D{
							{"$arrayElemAt",
								bson.A{
									"$channel",
									0,
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "Message"},
					{"as", "messages"},
					{"let", bson.D{{"channelIdentifier", "$channel.identifier"}}},
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
							bson.D{{"$addFields", bson.D{{"id", "$_id"}}}},
							bson.D{{"$limit", args.MessagesLimit}},
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
							bson.D{{"$addFields", bson.D{{"participant.id", "$participant._id"}}}},
						},
					},
				},
			},
		},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "Message"},
					{"as", "unreadCount"},
					{"let",
						bson.D{
							{"channelIdentifier", "$channel.identifier"},
							{"lastSeenMessageCreatedAt", "$lastSeenMessageCreatedAt"},
						},
					},
					{"pipeline",
						bson.A{
							bson.D{
								{"$match",
									bson.D{
										{"$expr",
											bson.D{
												{"$and",
													bson.A{
														bson.D{
															{"$eq",
																bson.A{
																	"$channel",
																	"$$channelIdentifier",
																},
															},
														},
														bson.D{
															{"$gt",
																bson.A{
																	"$createdAt",
																	"$$lastSeenMessageCreatedAt",
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
							bson.D{{"$count", "totalCount"}},
						},
					},
				},
			},
		},
		bson.D{
			{"$addFields",
				bson.D{
					{"unreadCount",
						bson.D{
							{"$arrayElemAt",
								bson.A{
									"$unreadCount",
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
					{"channel.id", "$channel._id"},
					{"channel.identifier", 1},
					{"channel.customData", 1},
					{"channel.status", 1},
					{"channel.createdAt", 1},
					{"messages.id", 1},
					{"messages.text", 1},
					{"messages.customData", 1},
					{"messages.type", 1},
					{"messages.subType", 1},
					{"messages.deletedAt", 1},
					{"messages.createdAt", 1},
					{"messages.participant.id", 1},
					{"messages.participant.identifier", 1},
					{"messages.participant.customData", 1},
					{"messages.participant.lastSeenAt", 1},
					{"messages.participant.isOnline", 1},
					{"unreadCount", "$unreadCount.totalCount"},
				},
			},
		},
	})

	var channelsDocuments []mongoSchemes.ChannelWithLastMessages

	err = channelsWithMsg.All(ctx, &channelsDocuments)

	if err != nil {
		return []mongoSchemes.ChannelWithLastMessages{}, err
	}

	return channelsDocuments, nil
}

type CountParticipantsInChannelArgs struct {
	ChannelIdentifier string
}
type CountParticipantsInChannelRes struct {
	ParticipantsCount       int64
	OnlineParticipantsCount int64
	Channel                 mongoSchemes.Channel
}

func CountParticipantsInChannel(args CountParticipantsInChannelArgs) (CountParticipantsInChannelRes, error) {
	var client = config.GetMongoClient()

	channelsWithMsg, err := client.Database(
		config.MainDatabase).Collection(mongoSchemes.ParticipantCollection).Aggregate(ctx, bson.A{
		bson.D{{"$match", bson.D{{"channel", args.ChannelIdentifier}}}},
		bson.D{
			{"$addFields",
				bson.D{
					{"participantNum", 1},
					{"onlineParticipantNum",
						bson.D{
							{"$cond",
								bson.A{
									bson.D{
										{"$eq",
											bson.A{
												"$isOnline",
												true,
											},
										},
									},
									1,
									0,
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{"$group",
				bson.D{
					{"_id", primitive.Null{}},
					{"channelIdentifier", bson.D{{"$first", "$channel"}}},
					{"participantsCount", bson.D{{"$sum", "$participantNum"}}},
					{"onlineParticipantsCount", bson.D{{"$sum", "$onlineParticipantNum"}}},
				},
			},
		},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "Channel"},
					{"localField", "channelIdentifier"},
					{"foreignField", "identifier"},
					{"as", "channel"},
				},
			},
		},
		bson.D{
			{"$addFields",
				bson.D{
					{"channel",
						bson.D{
							{"$arrayElemAt",
								bson.A{
									"$channel",
									0,
								},
							},
						},
					},
				},
			},
		},
	})

	var channelsDocuments []mongoSchemes.ChannelParticipantsCount

	err = channelsWithMsg.All(ctx, &channelsDocuments)

	if err != nil {
		return CountParticipantsInChannelRes{}, err
	}

	if len(channelsDocuments) != 1 {
		return CountParticipantsInChannelRes{}, nil
	}

	return CountParticipantsInChannelRes{
		ParticipantsCount:       channelsDocuments[0].ParticipantsCount,
		OnlineParticipantsCount: channelsDocuments[0].OnlineParticipantsCount,
		Channel:                 channelsDocuments[0].Channel,
	}, nil
}
