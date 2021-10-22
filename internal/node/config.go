package node

import (
	"context"
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"

	"github.com/kataras/golog"
	"github.com/pterm/pterm"
	api "github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/pkg/common"
)

// Error Definitions
var (
	logger             = golog.Child("internal/node")
	ErrEmptyQueue      = errors.New("No items in Transfer Queue.")
	ErrInvalidQuery    = errors.New("No SName or PeerID provided.")
	ErrMissingParam    = errors.New("Paramater is missing.")
	ErrProtocolsNotSet = errors.New("Node Protocol has not been initialized.")
)

// Option is a function that modifies the node options.
type Option func(*options)

// WithHost sets the host for RPC Stub Server
func WithHost(h string) Option {
	return func(o *options) {
		o.host = h
	}
}

// WithRequest sets the initialize request.
func WithRequest(req *api.InitializeRequest) Option {
	return func(o *options) {
		o.location = req.GetLocation()
		o.profile = req.GetProfile()
		o.connection = req.GetConnection()
	}
}

// WithMode starts the Client RPC server as a highway node.
func WithMode(m StubMode) Option {
	return func(o *options) {
		o.mode = m
	}
}

// WithPort sets the port for RPC Stub Server
func WithPort(p int) Option {
	return func(o *options) {
		o.port = p
	}
}

// options is a collection of options for the node.
type options struct {
	host       string
	connection common.Connection
	location   *common.Location
	mode       StubMode
	network    string
	port       int
	profile    *common.Profile
}

// defaultNodeOptions returns the default node options.
func defaultNodeOptions() *options {
	return &options{
		mode:       StubMode_LIB,
		location:   common.DefaultLocation(),
		connection: common.Connection_WIFI,
		network:    "tcp",
		host:       ":",
		port:       common.RPC_SERVER_PORT,
		profile:    common.NewDefaultProfile(),
	}
}

// Address returns the address of the node.
func (opts *options) Address() string {
	return fmt.Sprintf("%s%d", opts.host, opts.port)
}

// Apply applies Options to node
func (opts *options) Apply(ctx context.Context, node *Node) error {
	// Set Mode
	node.mode = opts.mode

	// Handle by Node Mode
	if opts.mode.HasClient() {
		logger.Debug("Starting Client stub...")
		// Client Node Type
		stub, err := node.startClientService(ctx, opts)
		if err != nil {
			logger.Error("Failed to start Client Service", err)
			return err
		}

		// Set Stub to node
		node.clientStub = stub

	} else {
		logger.Debug("Starting Highway stub...")
		// Highway Node Type
		stub, err := node.startHighwayService(ctx, opts)
		if err != nil {
			logger.Error("Failed to start Highway Service", err)
			return err
		}

		// Set Stub to node
		node.highwayStub = stub
	}
	return nil
}

// printTerminal is a helper function that prints to the terminal.
func (n *Node) printTerminal(title string, msg string) {
	if n.mode.IsCLI() {
		// Print a section with level one.
		pterm.DefaultSection.Println(title)
		// Print placeholder.
		pterm.Info.Println(msg)
	}
}

// promptTerminal is a helper function that prompts the user for input.
func (n *Node) promptTerminal(title string, onResult func(result bool)) error {
	if n.mode.IsCLI() {
		prompt := promptui.Prompt{
			Label:     title,
			IsConfirm: true,
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}
		onResult(result == "y")
	}
	return nil
}
