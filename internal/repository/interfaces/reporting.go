package interfaces

import (
	"boilerplate/internal/models"
	"context"
	"time"
)

// ReportRepository interface for report operations
type ReportRepository interface {
	Create(ctx context.Context, report *models.Report) error
	GetByID(ctx context.Context, id uint) (*models.Report, error)
	GetByName(ctx context.Context, name string) (*models.Report, error)
	Update(ctx context.Context, report *models.Report) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.Report, error)
	GetByUserID(ctx context.Context, userID uint) ([]*models.Report, error)
	GetByOutletID(ctx context.Context, outletID uint) ([]*models.Report, error)
	GetByType(ctx context.Context, reportType models.ReportTypeEnum) ([]*models.Report, error)
	GetByStatus(ctx context.Context, status models.ReportStatus) ([]*models.Report, error)
	GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.Report, error)
	UpdateStatus(ctx context.Context, id uint, status models.ReportStatus) error
	UpdateFilePath(ctx context.Context, id uint, filePath string) error
}

// PromotionRepository interface for promotion operations
type PromotionRepository interface {
	Create(ctx context.Context, promotion *models.Promotion) error
	GetByID(ctx context.Context, id uint) (*models.Promotion, error)
	GetByName(ctx context.Context, name string) (*models.Promotion, error)
	Update(ctx context.Context, promotion *models.Promotion) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.Promotion, error)
	GetActive(ctx context.Context) ([]*models.Promotion, error)
	GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.Promotion, error)
}