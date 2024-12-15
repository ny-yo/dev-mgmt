package database

import (
	"context"
	"dev-mgmt/internal/domain"
	"time"

	"gorm.io/gorm"
)

type Device struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Cert      string
	CreatedAt time.Time
}

type DeviceRepositoryImpl struct {
	db *gorm.DB
}

func NewDeviceRepositoryImpl(db *gorm.DB) *DeviceRepositoryImpl {
	return &DeviceRepositoryImpl{db: db}
}

func (r *DeviceRepositoryImpl) RegisterDevice(ctx context.Context, device *domain.Device) error {
	dbDevice := &Device{
		ID:        device.ID,
		Name:      device.Name,
		Cert:      device.Cert,
		CreatedAt: time.Now(),
	}
	return r.db.WithContext(ctx).Create(dbDevice).Error
}

func (r *DeviceRepositoryImpl) GetDeviceByCert(ctx context.Context, cert string) (*domain.Device, error) {
	var dbDevice Device
	if err := r.db.WithContext(ctx).Where("cert = ?", cert).First(&dbDevice).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrDeviceNotFound
		}
		return nil, err
	}

	return &domain.Device{
		ID:        dbDevice.ID,
		Name:      dbDevice.Name,
		Cert:      dbDevice.Cert,
		CreatedAt: dbDevice.CreatedAt.String(),
	}, nil
}
