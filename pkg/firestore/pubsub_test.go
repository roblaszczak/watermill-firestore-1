package firestore_test

import (
	"os"
	"testing"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/tests"

	"github.com/czeslavo/watermill-firestore/pkg/firestore"
)

func createPubSub(t *testing.T) (message.Publisher, message.Subscriber) {
	return createPubSubWithSubscriptionName(t, "topic")
}

func createPubSubWithSubscriptionName(t *testing.T, topic string) (message.Publisher, message.Subscriber) {
	logger := watermill.NewStdLogger(true, false)

	pub, err := firestore.NewPublisher(
		firestore.PublisherConfig{
			ProjectID: os.Getenv("FIRESTORE_PROJECT_ID"),
		},
		logger,
	)
	if err != nil {
		t.Fatal(err)
	}

	sub, err := firestore.NewSubscriber(
		firestore.SubscriberConfig{
			GenerateSubscriptionName: func(topic string) string {
				return topic + "_sub"
			},
			ProjectID: os.Getenv("FIRESTORE_PROJECT_ID"),
		},
		logger,
	)
	if err != nil {
		panic(err)
	}
	return pub, sub
}

func TestPublishSubscribe(t *testing.T) {
	tests.TestPubSub(
		t,
		tests.Features{
			ConsumerGroups:      true,
			ExactlyOnceDelivery: false,
			GuaranteedOrder:     false,
			Persistent:          true,
		},
		createPubSub,
		createPubSubWithSubscriptionName,
	)
}
