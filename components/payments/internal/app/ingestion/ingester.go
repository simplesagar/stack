package ingestion

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"

	"github.com/formancehq/payments/internal/app/models"
	"github.com/formancehq/stack/libs/go-libs/logging"
)

type Ingester interface {
	IngestPayments(ctx context.Context, batch PaymentBatch, commitState any) error
	IngestAccounts(ctx context.Context, batch AccountBatch) error
	GetTransfer(ctx context.Context, transferID uuid.UUID) (models.Transfer, error)
	UpdateTransferStatus(ctx context.Context, transferID uuid.UUID, status models.TransferStatus, reference, err string) error
}

type DefaultIngester struct {
	repo       Repository
	logger     logging.Logger
	provider   models.ConnectorProvider
	descriptor models.TaskDescriptor
	publisher  message.Publisher
}

type Repository interface {
	UpsertAccounts(ctx context.Context, provider models.ConnectorProvider, accounts []models.Account) error
	UpsertPayments(ctx context.Context, provider models.ConnectorProvider, payments []*models.Payment) error
	UpdateTaskState(ctx context.Context, provider models.ConnectorProvider, descriptor models.TaskDescriptor, state json.RawMessage) error

	GetTransfer(ctx context.Context, transferID uuid.UUID) (models.Transfer, error)
	UpdateTransferStatus(ctx context.Context, transferID uuid.UUID, status models.TransferStatus, reference, err string) error
}

func NewDefaultIngester(
	provider models.ConnectorProvider,
	descriptor models.TaskDescriptor,
	repo Repository,
	logger logging.Logger,
	publisher message.Publisher,
) *DefaultIngester {
	return &DefaultIngester{
		provider:   provider,
		descriptor: descriptor,
		repo:       repo,
		logger:     logger,
		publisher:  publisher,
	}
}