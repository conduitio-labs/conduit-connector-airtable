package source

import (
	"context"
	"github.com/AdamHaffar/conduit-connector-airtable/config"

	sdk "github.com/conduitio/conduit-connector-sdk"
)

type Source struct {
	sdk.UnimplementedSource

	config           config.Config
	lastPositionRead sdk.Position
}

func NewSource() sdk.Source {
	return &Source{}
}

func (s *Source) Configure(ctx context.Context, cfg map[string]string) error {
	sdk.Logger(ctx).Info().Msg("Configuring a Source Connector...")
	SourceConfig, err := config.ParseBaseConfig(cfg)
	if err != nil {
		return err
	}

	s.config = SourceConfig
	return nil
}

func (s *Source) Open(ctx context.Context, pos sdk.Position) error {
	return nil
}

func (s *Source) Read(ctx context.Context) (sdk.Record, error) {
	return sdk.Record{}, nil
}

func (s *Source) Ack(ctx context.Context, position sdk.Position) error {
	return nil
}

func (s *Source) Teardown(ctx context.Context) error {
	return nil
}
