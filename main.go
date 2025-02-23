// main.go - compile with -buildmode=plugin
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/zeromq/goczmq"

	// Import the core plugin API definitions.
	"github.com/withObsrvr/pluginapi"
)

// SaveToZeroMQ implements pluginapi.Consumer.
type SaveToZeroMQ struct {
	publisher *goczmq.Sock
	address   string
}

// NewSaveToZeroMQ creates a new instance of SaveToZeroMQ based on the provided configuration.
func NewSaveToZeroMQ(config map[string]interface{}) (*SaveToZeroMQ, error) {
	address, ok := config["address"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid configuration for SaveToZeroMQ: missing 'address'")
	}

	publisher, err := goczmq.NewPub(address)
	if err != nil {
		return nil, fmt.Errorf("error creating ZeroMQ publisher: %w", err)
	}

	return &SaveToZeroMQ{
		publisher: publisher,
		address:   address,
	}, nil
}

// Process implements the pluginapi.Consumer interface.
// It expects the message payload to be a []byte and writes it to the ZeroMQ publisher.
func (z *SaveToZeroMQ) Process(ctx context.Context, msg pluginapi.Message) error {
	log.Printf("SaveToZeroMQ: Processing message")
	payload, ok := msg.Payload.([]byte)
	if !ok {
		return fmt.Errorf("expected []byte type for payload, got %T", msg.Payload)
	}

	_, err := z.publisher.Write(payload)
	if err != nil {
		return fmt.Errorf("error writing to ZeroMQ: %w", err)
	}

	return nil
}

// Close shuts down the ZeroMQ publisher.
func (z *SaveToZeroMQ) Close() error {
	z.publisher.Destroy()
	return nil
}

// Name returns the plugin name.
func (z *SaveToZeroMQ) Name() string {
	return "SaveToZeroMQ"
}

// Version returns the plugin version.
func (z *SaveToZeroMQ) Version() string {
	return "1.0.0"
}

// Type returns the plugin type (Consumer in this case).
func (z *SaveToZeroMQ) Type() pluginapi.PluginType {
	return pluginapi.ConsumerPlugin
}

// Initialize sets up the plugin with configuration.
// In this implementation, we delegate to NewSaveToZeroMQ and copy its values.
func (z *SaveToZeroMQ) Initialize(config map[string]interface{}) error {
	instance, err := NewSaveToZeroMQ(config)
	if err != nil {
		return err
	}
	*z = *instance
	return nil
}

// Exported New function allows dynamic loading by the plugin manager.
// It returns a new instance that satisfies pluginapi.Plugin.
func New() pluginapi.Plugin {
	return &SaveToZeroMQ{}
}
